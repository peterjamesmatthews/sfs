package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/graph/model"
	"pjm.dev/sfs/memdb"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := newSeededDatabase()
	http.Handle("/graphql", graph.GetGQLHandler(&db, &db, &db))
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func newSeededDatabase() memdb.MemDatabase {
	peter := &model.User{ID: uuid.NewString(), Name: "Peter"}
	users := []*model.User{peter}

	root := &model.Folder{}

	foo := &model.Folder{
		ID:     uuid.NewString(),
		Name:   "Foo",
		Owner:  peter,
		Parent: root,
	}

	bar := &model.File{
		ID:      uuid.NewString(),
		Name:    "Bar",
		Owner:   peter,
		Parent:  root,
		Content: "Hello World!",
	}

	root.Children = []model.Node{foo, bar}

	return memdb.MemDatabase{Root: root, Users: users}
}
