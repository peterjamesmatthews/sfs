package graph

import (
	"context"
	"errors"
	"fmt"
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

func handleGettingUserFromContext(ctx context.Context, authN Authenticator) (User, error) {
	user, err := authN.FromContext(ctx)
	if errors.Is(err, ErrUnauthorized) {
		return User{}, err
	} else if err != nil {
		return User{}, fmt.Errorf("failed to get authenticated user: %w", err)
	}
	return user, nil
}
