package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

func GetGQLHandler(authN Authenticator, sfs SharedFileSystemer, uuidGen UUIDGenerator) http.Handler {
	var gqlHandler http.Handler

	resolver := Resolver{
		AuthN:   authN,
		SFS:     sfs,
		UUIDGen: uuidGen,
	}

	config := Config{Resolvers: &resolver}

	executableSchema := NewExecutableSchema(config)

	gqlHandler = handler.NewDefaultServer(executableSchema)

	gqlHandler = authN.WrapInAuthentication(gqlHandler)

	return gqlHandler
}
