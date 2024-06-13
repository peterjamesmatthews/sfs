package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"pjm.dev/sfs/app"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/server"
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
		log.Fatalf("failed to initialize db: %v", err)
	}

	// initialize app
	app := app.New(config.App, db)

	// initialize graph
	graphHandler := graph.New(graph.Resolver{SharedFileSystem: &app})

	// initialize server
	handler := server.New(config.Server, graphHandler)

	// start server
	log.Fatalln(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", config.Server.Hostname, config.Server.Port),
			handler,
		),
	)
}
