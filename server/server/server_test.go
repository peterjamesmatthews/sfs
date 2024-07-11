package server_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	m "github.com/stretchr/testify/mock"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/meta"
	"pjm.dev/sfs/server"
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

// mock provides a mockable interface for many of the project's interfaces.
type mock struct {
	m.Mock
}

func newTestServer(t *testing.T) (*httptest.Server, *mock, *pgx.Conn) {
	t.Helper()

	// initialize postgres container for database
	dbContainer := newPostgresContainer(t)

	// configure stack's config to use the database container
	cfg := getConfigFromDatabaseContainer(dbContainer, t)

	// initialize mock for stack
	mock := new(mock)

	// initialize stack
	stack, err := config.NewStack(cfg, config.WithAuth0er(mock))
	if err != nil {
		t.Fatalf("failed to initialize test server: %v", err)
	}

	// return test server, mock, and database connection
	return httptest.NewServer(stack.Server), mock, stack.Database
}

var cfg = config.Config{
	Database: db.Config{
		User:     "foo",
		Password: "bar",
		Name:     "baz",
	},
}

// getConfigFromDatabaseContainer returns a config.Config with the database connection details from a postgres container.
func getConfigFromDatabaseContainer(dbContainer *postgres.PostgresContainer, t *testing.T) config.Config {
	connectionString, err := dbContainer.ConnectionString(context.Background())
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	connectionURL, err := url.Parse(connectionString)
	if err != nil {
		t.Fatalf("failed to parse connection URL: %v", err)
	}

	port, err := strconv.Atoi(connectionURL.Port())
	if err != nil {
		t.Fatalf("failed to parse connection port: %v", err)
	}

	cfg := cfg.Clone()
	cfg.Database.Hostname = connectionURL.Hostname()
	cfg.Database.Port = port
	cfg.Server = server.Config{GraphEndpoint: "graph"}
	return cfg
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
		postgres.WithUsername(cfg.Database.User),
		postgres.WithPassword(cfg.Database.Password),
		postgres.WithDatabase(cfg.Database.Name),
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

// dumpDatabase returns a string result of the pg_dump command on a given database.
//
// If an error occurs, this function calls t.Fatalf with the error message.
func dumpDatabase(db *pgx.Conn, t *testing.T) string {
	cmd := exec.Command(
		"pg_dump",
		"-d", db.Config().Database,
		"-h", db.Config().Host,
		"-p", fmt.Sprintf("%d", db.Config().Port),
		"-U", db.Config().User,
	)
	cmd.Env = append(cmd.Environ(), "PGPASSWORD="+db.Config().Password)
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to dump database: %v", err)
	}
	return string(output)
}

// TestNewPostgresContainer simply tests the newPostgresContainer function executes without error.
func TestNewPostgresContainer(t *testing.T) {
	newPostgresContainer(t)
}
