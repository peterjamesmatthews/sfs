package main

import (
	"context"
	"encoding/json"
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
	// init context
	ctx := context.Background()

	// init logging
	log.Default().SetFlags(0)

	// init config
	config, err := config.New(ctx)
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	// pretty print config
	configBytes, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		log.Fatalf("failed to marshal config: %v", err)
	}
	log.Printf("initializing with config: %v", string(configBytes))

	// init db from config
	db, err := db.New(config.Database)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// init app
	app := app.New(db)
	gqlHandler := graph.GetGQLHandler(app, app, app)

	// init server from config
	pattern := fmt.Sprintf("/%s", config.Server.GraphEndpoint)
	gqlHandler = WrapHandler(gqlHandler, &LoggingHandler{}, &CORSHandler{})
	http.Handle(pattern, gqlHandler)
	http.Handle("/", playground.Handler("GraphQL playground", pattern))

	// start server
	log.Printf("serving GraphQL at http://%s:%s%s", config.Server.Hostname, config.Server.Port, pattern)
	log.Fatal(http.ListenAndServe(config.Server.Hostname+":"+config.Server.Port, nil))
}
