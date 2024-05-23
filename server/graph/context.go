package graph

import (
	"context"
	"errors"
	"log"
)

type contextKey string

var authorizationKey = contextKey("authorization")

func getContextWithAuthorization(ctx context.Context, authorization string) context.Context {
	return context.WithValue(ctx, authorizationKey, authorization)
}

func getAuthorizationFromContext(ctx context.Context) string {
	value := ctx.Value(authorizationKey)
	if value == nil {
		log.Printf("authorization not found in context")
		return ""
	}

	authorization, ok := value.(string)
	if !ok {
		log.Printf("invalid authorization in context: %v", value)
		return ""
	}

	return authorization
}

var resolverKey = contextKey("resolver")

func getContextWithResolver(ctx context.Context, resolver Resolver) context.Context {
	return context.WithValue(ctx, resolverKey, resolver)
}

func getResolverFromContext(ctx context.Context) (Resolver, error) {
	value := ctx.Value(resolverKey)
	if value == nil {
		log.Printf("resolver not found in context")
		return Resolver{}, errors.New("resolver not found in context")
	}

	resolver, ok := value.(Resolver)
	if !ok {
		log.Printf("invalid resolver in context: %v", value)
		return Resolver{}, errors.New("invalid resolver in context")
	}

	return resolver, nil
}

var userKey = contextKey("user")

func getContextWithUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func getUserFromContext(ctx context.Context) (User, error) {
	value := ctx.Value(userKey)
	if value == nil {
		log.Printf("user not found in context")
		return User{}, errors.New("user not found in context")
	}

	user, ok := value.(User)
	if !ok {
		log.Printf("invalid user in context: %v", value)
		return User{}, errors.New("invalid user in context")
	}

	return user, nil
}
