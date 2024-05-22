package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"pjm.dev/sfs/config"
)

func New(config config.ServerConfig, resolver Resolver) http.Handler {
	executableSchema := NewExecutableSchema(Config{Resolvers: &resolver})
	return handler.NewDefaultServer(executableSchema)
}
