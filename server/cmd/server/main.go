package main

import (
	"fmt"
	"log"
	"net/http"

	"pjm.dev/sfs/app"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/server"
)

func main() {
	// initialize logging
	log.Default().SetFlags(0)

	// initialize config
	config, err := config.New()
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

	// initialize handler
	handler := server.New(config.Server, app)

	// start server
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Server.Hostname, config.Server.Port), handler))
}
