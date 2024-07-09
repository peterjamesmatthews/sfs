package config

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/auth0"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/server"
)

type Stack struct {
	Database *pgx.Conn
	App      app.App
	Server   http.Handler
}

// NewStack initializes a tech stack from a config.
//
// The stack includes:
//   - A Postgres database connection
//   - An application layer
//   - A GraphQL server's handler
func NewStack(config Config, auth0er app.Auth0er) (Stack, error) {
	// initialize db
	db, err := db.New(config.Database)
	if err != nil {
		return Stack{}, fmt.Errorf("failed to initialize db: %w", err)
	}

	// initialize auth0
	auth0 := auth0.New(config.Auth0)

	// initialize app
	app := app.New(config.App, db, auth0)

	// TODO ref to allow optional components of stack in a more idiomatic way
	if auth0er != nil {
		app.SetAuth0er(auth0er)
	}

	// initialize graph
	graphHandler := graph.New(graph.Resolver{SharedFileSystem: &app})

	server := server.New(config.Server, graphHandler)

	// initialize server
	return Stack{Database: db, App: app, Server: server}, nil
}
