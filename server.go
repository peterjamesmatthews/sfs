package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/memdb"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var gqlHandler http.Handler
	db := memdb.NewSeededDatabase()
	resolver := graph.Resolver{AuthN: &db, SFS: &db}
	gqlHandler = handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))
	gqlHandler = db.WrapInAuthentication(gqlHandler)

	http.Handle("/graphql", gqlHandler)
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
