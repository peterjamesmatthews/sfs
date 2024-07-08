package server

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)

func New(config Config, graph http.Handler) http.Handler {
	pattern := fmt.Sprintf("/%s", config.GraphEndpoint)
	mux := http.NewServeMux()

	// register graph's API
	mux.Handle(pattern, graph)

	// register graph's playground
	mux.Handle("/", playground.Handler("SFS Playground", pattern))

	return mux
}
