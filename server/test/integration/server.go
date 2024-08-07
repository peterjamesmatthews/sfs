package integration

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	m "github.com/stretchr/testify/mock"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/meta"
)

// test represents a single test for server_test tests.
type test struct {
	// identifier for the test
	name string

	// database state before the test
	seed *os.File

	// request to send to the server
	request *http.Request

	// response to expect from the server
	response *http.Response

	// database state after the test
	dump *os.File
}

// stackMock provides a mock for this stack's interfaces
type stackMock struct {
	m.Mock
}

func newTestServer(t *testing.T) (*httptest.Server, *stackMock, config.Stack) {
	t.Helper()

	// initialize postgres container for database
	postgresContainer := newPostgresContainer(t)

	// configure stack's config to use the database container
	cfg := config.Config{
		Database: getDatabaseConfigFromContainer(t, postgresContainer),
	}

	// initialize mockable for stack
	mockable := new(stackMock)

	// initialize stack
	stack, err := config.NewStack(cfg, config.WithAuth0er(mockable))
	if err != nil {
		t.Fatalf("failed to initialize test server: %v", err)
	}

	// return test server, mock, and database connection
	return httptest.NewServer(stack.Server), mockable, stack
}

// getDatabaseConfigFromContainer returns a config.Config with the database connection details from a postgres container.
func getDatabaseConfigFromContainer(t *testing.T, dbContainer *postgres.PostgresContainer) db.Config {
	connectionString, err := dbContainer.ConnectionString(context.Background())
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	url, err := url.Parse(connectionString)
	if err != nil {
		t.Fatalf("failed to parse connection URL: %v", err)
	}

	_, err = strconv.Atoi(url.Port())
	if err != nil {
		t.Fatalf("failed to parse connection port: %v", err)
	}

	port, err := strconv.Atoi(url.Port())
	if err != nil {
		t.Fatalf("failed to parse connection port: %v", err)
	}

	password, set := url.User.Password()
	if !set {
		t.Fatalf("failed to parse unset connection password: %v", err)
	}

	return db.Config{
		User:     url.User.Username(),
		Password: password,
		Hostname: url.Hostname(),
		Port:     port,
		Name:     strings.TrimPrefix(url.Path, "/"),
	}
}

// newPostgresContainer starts a new postgres container with the project's database schema.
func newPostgresContainer(t *testing.T) *postgres.PostgresContainer {
	t.Helper()

	ctx := context.Background()

	// find database schema migration files; they should be alphabetically sorted
	migrations, err := filepath.Glob(filepath.Join(meta.Root, "db", "migrations", "*.sql"))
	if err != nil {
		t.Fatalf("failed to find migration files: %s", err)
	}

	// start postgres container with database schema migrations
	container, err := postgres.Run(ctx,
		"docker.io/postgres:16-bullseye",
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
		postgres.WithDatabase("postgres"),
		postgres.WithInitScripts(migrations...),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}

	// register container termination on test cleanup
	t.Cleanup(func() {
		container.Terminate(ctx)
	})

	return container
}
