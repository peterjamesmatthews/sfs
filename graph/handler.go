package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

func GetGQLHandler(authN Authenticator, sfs SharedFileSystemer) http.Handler {
	var gqlHandler http.Handler
	resolver := Resolver{AuthN: authN, SFS: sfs}
	config := Config{Resolvers: &resolver}
	executableSchema := NewExecutableSchema(config)
	gqlHandler = handler.NewDefaultServer(executableSchema)
	gqlHandler = authN.WrapInAuthentication(gqlHandler)
	return gqlHandler
}
