package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
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

type mock struct {
	m.Mock
}

var cfg = config.Config{
	Database: db.Config{
		User:     "foo",
		Password: "bar",
		Name:     "baz",
	},
}

func newTestServer(t *testing.T) (*httptest.Server, *mock, *pgx.Conn) {
	t.Helper()

	dbContainer := newPostgresContainer(t)

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

	mock := new(mock)

	stack, err := config.NewStack(cfg, mock)
	if err != nil {
		t.Fatalf("failed to initialize test server: %v", err)
	}

	return httptest.NewServer(stack.Server), mock, stack.Database
}

func newPostgresContainer(t *testing.T) *postgres.PostgresContainer {
	t.Helper()

	ctx := context.Background()

	migrations, err := filepath.Glob(filepath.Join(meta.Root, "db", "migrations", "*.sql"))
	if err != nil {
		t.Fatalf("failed to find migration files: %s", err)
	}

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

	return container
}

func TestNewPostgresContainer(t *testing.T) {
	container := newPostgresContainer(t)
	t.Cleanup(func() {
		container.Terminate(context.Background())
	})
}
