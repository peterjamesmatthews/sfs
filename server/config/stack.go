package config

import (
	"database/sql"
	"fmt"
	"net/http"

	"pjm.dev/sfs/app"
	"pjm.dev/sfs/auth0"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/server"
)

type Stack struct {
	App      app.App
	Database *sql.DB
	Server   http.Handler
}

type StackOption func(*Stack)

// WithAuth0er overrides the default Auth0er in the stack.
func WithAuth0er(auth0er app.Auth0er) StackOption {
	return func(s *Stack) {
		s.App.SetAuth0er(auth0er)
	}
}

// NewStack initializes a tech stack from a config.
//
// The stack includes:
//   - A Postgres database connection
//   - An application layer
//   - A GraphQL server's handler
func NewStack(config Config, stackOptions ...StackOption) (Stack, error) {
	// initialize stack
	var stack Stack

	// initialize db
	db, err := db.New(config.Database)
	if err != nil {
		return Stack{}, fmt.Errorf("failed to initialize db: %w", err)
	}

	// set db in stack
	stack.Database = db

	// initialize auth0
	auth0 := auth0.New(config.Auth0)

	// initialize app
	app := app.New(config.App, db, auth0)

	// set app in stack
	stack.App = app

	// apply stack options to stack
	for _, stackOption := range stackOptions {
		stackOption(&stack)
	}

	// initialize graph
	graphHandler := graph.New(graph.Resolver{SharedFileSystem: &stack.App})

	// initialize server
	server := server.New(config.Server, graphHandler)

	// set server in stack
	stack.Server = server

	// return stack
	return stack, nil
}
