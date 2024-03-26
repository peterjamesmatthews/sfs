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
	// init root user (owner of root folder)
	rootUser := &graph.User{ID: uuid.NewString(), Name: mem.RootName}

	// init users
	amos := &graph.User{ID: uuid.NewString(), Name: "Amos"}
	matthew := &graph.User{ID: uuid.NewString(), Name: "Matthew"}
	nick := &graph.User{ID: uuid.NewString(), Name: "Nick"}
	users := []*graph.User{amos, matthew, nick}

	// init root
	root := &graph.Folder{
		ID:    uuid.NewString(),
		Name:  mem.RootName,
		Owner: rootUser,
	}

	// init access
	access := []*graph.Access{}

	// grant all users read access to root
	for _, user := range users {
		access = append(access, &graph.Access{User: user, Type: graph.AccessTypeRead, Target: root})
	}

	// init root's children
	root.Children = []graph.Node{}

	// add matthew's nodes
	root.Children = append(root.Children,
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
	)

	// add nick's nodes
	root.Children = append(root.Children,
		&graph.File{
			ID:      uuid.NewString(),
			Name:    "Passwords",
			Owner:   nick,
			Parent:  root,
			Content: "nick-is-cool",
		},
	)

	// give all node owners read access to their nodes
	for _, node := range root.Children {
		access = append(access, &graph.Access{User: node.GetOwner(), Type: graph.AccessTypeRead, Target: node})
	}

	// give amos read access to all nodes
	for _, node := range root.Children {
		access = append(access, &graph.Access{User: amos, Type: graph.AccessTypeRead, Target: node})
	}

	return mem.Database{Root: root, Users: users, Access: access}
}
