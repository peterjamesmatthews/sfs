package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"pjm.dev/sfs/config"
)

func NewHandler(config config.ServerConfig, resolver Resolver) http.Handler {
	executableSchema := NewExecutableSchema(Config{Resolvers: &resolver})
	handler := handler.NewDefaultServer(executableSchema)
	return resolver.AuthN.WrapInAuthentication(handler)
}
