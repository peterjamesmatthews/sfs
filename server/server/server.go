package server

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)

func New(config Config, graphHandler http.Handler) http.Handler {
	mux := http.NewServeMux()
	registerGraphHandlers(config, mux, graphHandler)
	return mux
}

func registerGraphHandlers(config Config, mux *http.ServeMux, graphHandler http.Handler) {
	graphPattern := fmt.Sprintf("/%s", config.GraphEndpoint)

	// register graph's handler
	mux.Handle(graphPattern, graphHandler)

	// register graph's playground handler
	mux.Handle("/", playground.Handler("SFS Playground", graphPattern))
}
