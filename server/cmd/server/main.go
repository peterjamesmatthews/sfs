package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/graph/model"
	"pjm.dev/sfs/mem"
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

func newSeededDatabase() mem.Database {
	matthew := &model.User{ID: uuid.NewString(), Name: "Matthew"}
	nick := &model.User{ID: uuid.NewString(), Name: "Nick"}
	users := []*model.User{matthew, nick}

	root := &model.Folder{}
	root.Children = []model.Node{
		&model.Folder{
			ID:     uuid.NewString(),
			Name:   "Empty Folder",
			Owner:  matthew,
			Parent: root,
		},
		&model.File{
			ID:      uuid.NewString(),
			Name:    "Greeting",
			Owner:   matthew,
			Parent:  root,
			Content: "Hello World!",
		},
		&model.File{
			ID:      uuid.NewString(),
			Name:    "Passwords",
			Owner:   nick,
			Parent:  root,
			Content: "nick-is-cool",
		},
	}

	access := []*model.Access{
		{User: matthew, Type: model.AccessTypeRead, Target: root},
		{User: nick, Type: model.AccessTypeRead, Target: root},
	}

	return mem.Database{Root: root, Users: users, Access: access}
}
