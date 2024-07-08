package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/meta"
	"pjm.dev/sfs/server"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	// TODO cache this container across tests; test independence should be solved with transactions
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

	cfg := config.Config{}

	cfg.Database = db.Config{
		Hostname: connectionURL.Hostname(),
		Port:     port,
		Name:     "postgres",
		Password: "password",
		User:     "postgres",
	}

	cfg.Server = server.Config{GraphEndpoint: "graph"}

	db, app, server, err := config.NewStack(cfg)
	if err != nil {
		t.Fatalf("failed to initialize test server: %v", err)
	}

	// wrap each request in a transaction
	server = wrapInTransaction(server, db, t, app)

	return httptest.NewServer(server)
}

func wrapInTransaction(server http.Handler, db *pgx.Conn, t *testing.T, app app.App) http.Handler {
	server = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.Begin(r.Context())
		if err != nil {
			t.Fatalf("failed to begin transaction: %v", err)
		}

		app.SetDatabase(tx.Conn())

		server.ServeHTTP(w, r)

		if err := tx.Rollback(r.Context()); err != nil {
			t.Fatalf("failed to rollback transaction: %v", err)
		}
	})
	return server
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
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("password"),
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
