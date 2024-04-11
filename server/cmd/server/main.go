package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/graph"
)

func main() {
	// initialize context
	ctx := context.Background()

	// initialize logging
	log.Default().SetFlags(0)

	// initialize config
	config, err := config.New(ctx)
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	log.Printf("initializing with config: %s", config)

	// initialize db
	db, err := db.New(config.Database)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// initialize app
	app := app.New(db)
	gqlHandler := graph.GetGQLHandler(app, app, app)

	// initialize server
	pattern := fmt.Sprintf("/%s", config.Server.GraphEndpoint)
	gqlHandler = WrapHandler(gqlHandler, &LoggingHandler{}, &CORSHandler{})
	http.Handle(pattern, gqlHandler)
	http.Handle("/", playground.Handler("GraphQL playground", pattern))

	// start server
	log.Printf("serving GraphQL at http://%s:%s%s", config.Server.Hostname, config.Server.Port, pattern)
	log.Fatal(http.ListenAndServe(config.Server.Hostname+":"+config.Server.Port, nil))
}
