package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/sethvargo/go-envconfig"
	"pjm.dev/sfs/db"
	"pjm.dev/sfs/env"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/mem"
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
	_, err = db.Initialize(config.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// // init server from config
	// pattern := fmt.Sprintf("/%s", config.Server.GraphEndpoint)
	// gqlHandler := graph.GetGQLHandler(&db, &db, &db)
	// gqlHandler = WrapHandler(gqlHandler, &LoggingHandler{}, &CORSHandler{})
	// http.Handle(pattern, gqlHandler)
	// http.Handle("/", playground.Handler("GraphQL playground", pattern))

	// // start serving
	// log.Printf("serving GraphQL at http://%s:%s%s", config.Server.Hostname, config.Server.Port, pattern)
	// log.Fatal(http.ListenAndServe(config.Server.Hostname+":"+config.Server.Port, nil))
}

func newSeededDatabase() mem.Database {
	// init root user (owner of root folder)
	rootUser := &graph.User{ID: uuid.NewString(), Name: mem.RootName}

	// init users
	amos := &graph.User{ID: uuid.NewString(), Name: "Amos"}
	jack := &graph.User{ID: uuid.NewString(), Name: "Jack"}
	nick := &graph.User{ID: uuid.NewString(), Name: "Nick"}
	users := []*graph.User{amos, jack, nick}

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
			Name:   "EmptyFolder",
			Owner:  jack,
			Parent: root,
		},
		&graph.File{
			ID:      uuid.NewString(),
			Name:    "Greeting",
			Owner:   jack,
			Parent:  root,
			Content: "Hello World!",
		},
	)

	// add nick's nodes
	nicksPasswords := &graph.Folder{
		ID:     uuid.NewString(),
		Name:   "FolderForPasswords",
		Owner:  nick,
		Parent: root,
	}

	root.Children = append(root.Children, nicksPasswords)

	nicksWarning := &graph.Folder{
		ID:     uuid.NewString(),
		Name:   "WARNING_SUPER_SECRET_PASSWORDS",
		Owner:  nick,
		Parent: nicksPasswords,
	}

	nicksPasswords.Children = append(nicksPasswords.Children, nicksWarning)

	passwords := &graph.File{
		ID:      uuid.NewString(),
		Name:    "Passwords",
		Owner:   nick,
		Parent:  nicksWarning,
		Content: "nick-is-super-cool",
	}

	nicksWarning.Children = append(nicksWarning.Children, passwords)

	// give all node owners read access to their nodes
	for _, node := range root.Children {
		access = append(access, &graph.Access{User: node.GetOwner(), Type: graph.AccessTypeRead, Target: node})
	}

	// give nick read access to all of his deep nodes
	nicksDeepNodes := []graph.Node{nicksWarning, passwords}
	for _, node := range nicksDeepNodes {
		access = append(access, &graph.Access{User: nick, Type: graph.AccessTypeRead, Target: node})
	}

	// give amos read access to all nodes
	for _, node := range root.Children {
		access = append(access, &graph.Access{User: amos, Type: graph.AccessTypeRead, Target: node})
	}

	return mem.Database{Root: root, Users: users, Access: access}
}
