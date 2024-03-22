package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/mem"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := newSeededDatabase()

	gqlHandler := graph.GetGQLHandler(&db, &db, &db)
	gqlHandler = WrapHandler(gqlHandler, &LoggingHandler{}, &CORSHandler{})

	http.Handle("/graphql", gqlHandler)
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	log.Printf("serving http://localhost:%s/graphql", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func newSeededDatabase() mem.Database {
	matthew := &graph.User{ID: uuid.NewString(), Name: "Matthew"}
	nick := &graph.User{ID: uuid.NewString(), Name: "Nick"}
	users := []*graph.User{matthew, nick}

	root := &graph.Folder{}
	root.Children = []graph.Node{
		&graph.Folder{
			ID:     uuid.NewString(),
			Name:   "Empty Folder",
			Owner:  matthew,
			Parent: root,
		},
		&graph.File{
			ID:      uuid.NewString(),
			Name:    "Greeting",
			Owner:   matthew,
			Parent:  root,
			Content: "Hello World!",
		},
		&graph.File{
			ID:      uuid.NewString(),
			Name:    "Passwords",
			Owner:   nick,
			Parent:  root,
			Content: "nick-is-cool",
		},
	}

	access := []*graph.Access{
		{User: matthew, Type: graph.AccessTypeRead, Target: root},
		{User: nick, Type: graph.AccessTypeRead, Target: root},
	}

	return mem.Database{Root: root, Users: users, Access: access}
}
