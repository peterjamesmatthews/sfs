package config

import (
	"fmt"
	"net/http"

	"pjm.dev/sfs/app"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/server"
)

func NewHandler(config Config) (http.Handler, error) {
	// initialize db
	db, err := db.New(config.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize db: %w", err)
	}

	// initialize app
	app := app.New(config.App, db)

	// initialize graph
	graphHandler := graph.New(graph.Resolver{SharedFileSystem: &app})

	// initialize server
	return server.New(config.Server, graphHandler), nil
}
