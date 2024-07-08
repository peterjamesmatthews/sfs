package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"pjm.dev/sfs/config"
)

func main() {
	// initialize logging
	log.SetFlags(0)
	log.Println("Hello from pjm.dev/sfs/cmd/server!")

	// initialize config
	cfg, err := config.New(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	log.Printf("initializing with config: %s", cfg)

	// initialize handler
	_, _, handler, err := config.NewStack(cfg)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	// construct server's address
	addr := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)
	log.Printf("server initialized, listening on %s", addr)

	// start server
	log.Fatalln(http.ListenAndServe(addr, handler))
}
