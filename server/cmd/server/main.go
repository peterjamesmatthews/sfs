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
	// initialize logging
	log.SetFlags(0)
	log.Println("Hello from pjm.dev/sfs/cmd/server!")

	// initialize config
	config, err := config.New(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	log.Printf("initializing with config: %s\n", config)

	// initialize db
	db, err := db.New(config.Database)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// initialize app
	app := app.New(db)

	// initialize graph
	graphHandler := graph.NewHandler(config.Server, graph.Resolver{
		SFS:     &app,
		AuthN:   &app,
		UUIDGen: &app,
	})

	// initialize server
	mux := http.NewServeMux()

	// register graph handler
	graphPattern := fmt.Sprintf("/%s", config.Server.GraphEndpoint)
	mux.Handle(graphPattern, graphHandler)

	// register graph's playground handler
	mux.Handle("/", playground.Handler("SFS Playground", graphPattern))

	// start server
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Server.Hostname, config.Server.Port), mux))
}
