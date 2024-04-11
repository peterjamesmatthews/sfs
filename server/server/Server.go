package server

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"pjm.dev/sfs/app"
	"pjm.dev/sfs/config"
	"pjm.dev/sfs/graph"
)

type Server interface {
	Serve() error
}

func New(config config.ServerConfig, app app.Apper) Server {
	mux := http.NewServeMux()

	// initialize graph handler
	graphPattern := fmt.Sprintf("/%s", config.GraphEndpoint)
	graphHandler := graph.NewHandler(app, app, app)
	graphHandler = wrapHandler(
		graphHandler,
		&loggingHandler{},
		&corsHandler{},
	)
	mux.Handle(graphPattern, graphHandler)

	// initialize playground handler
	playgroundPattern := "/"
	playgroundHandler := playground.Handler("SFS Playground", graphPattern)
	mux.Handle(playgroundPattern, playgroundHandler)

	return &server{
		addr:    fmt.Sprintf("%s:%s", config.Hostname, config.Port),
		handler: mux,
	}
}

type server struct {
	addr    string
	handler http.Handler
}

func (s *server) Serve() error {
	return http.ListenAndServe(s.addr, s.handler)
}
