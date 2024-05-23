package graph

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

// authenticated implements the @Authenticated directive by resolving an Authorization header to a User.
//
// # Errors
//   - ErrUnauthorized: if the Authorization header is missing or invalid.
func authenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	// get authorization from context
	auth := getAuthorizationFromContext(ctx)
	if auth == "" {
		return nil, ErrUnauthorized
	}

	// get resolver from context
	resolver, err := getResolverFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get resolver from context: %w", err)
	}

	// authenticate user from authorization
	user, err := resolver.SharedFileSystem.Authenticate(auth)
	if errors.Is(err, ErrUnauthorized) {
		return nil, ErrUnauthorized
	} else if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %w", err)
	}

	// set user in context
	ctx = getContextWithUser(ctx, user)

	// return next resolver
	return next(ctx)
}
