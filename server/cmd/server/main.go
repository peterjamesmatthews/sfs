package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sethvargo/go-envconfig"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/env"
	"pjm.dev/sfs/graph"
)

func main() {
	// init context
	ctx := context.Background()

	// init logging
	log.Default().SetFlags(0)

	// init config
	var config env.Config
	err := envconfig.Process(ctx, &config)
	if err != nil {
		log.Fatalf("failed to process config from environment: %v", err)
	}

	// pretty print config
	configBytes, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		log.Fatalf("failed to marshal config: %v", err)
	}
	log.Printf("initializing with config: %v", string(configBytes))

	// init db from config
	db, err := db.Initialize(config.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
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
