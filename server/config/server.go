package config

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/server"
)

// NewStack initializes a tech stack from a config.
//
// The stack includes:
//   - A Postgres database connection
//   - An application layer
//   - A GraphQL server's handler
func NewStack(config Config) (*pgx.Conn, app.App, http.Handler, error) {
	// initialize db
	db, err := db.New(config.Database)
	if err != nil {
		return nil, app.App{}, nil, fmt.Errorf("failed to initialize db: %w", err)
	}

	// initialize app
	app := app.New(config.App, db)

	// initialize graph
	graphHandler := graph.New(graph.Resolver{SharedFileSystem: &app})

	// initialize server
	return db, app, server.New(config.Server, graphHandler), nil
}
