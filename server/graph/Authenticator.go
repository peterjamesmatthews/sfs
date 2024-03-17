package graph

import (
	"context"
	"net/http"

	"pjm.dev/sfs/graph/model"
)

type Authenticator interface {
	// Authenticate determines the requesting user.
	//
	// # Arguments
	//  - r: The http request to authenticate.
	//
	// # Returns
	//  - The user who is making the request.
	//
	// # Errors
	//  - `ErrUnauthorized` if an authenticated user cannot be determined.
	Authenticate(*http.Request) (model.User, error)

	// WithUser wraps a user in a context.
	//
	// # Arguments
	//  - ctx: The context to wrap the user in.
	//  - user: The user to wrap in the context.
	//
	// # Returns
	//  - A new context with the user wrapped in it.
	WithUser(context.Context, model.User) context.Context

	// FromContext extracts a user from a context.
	//
	// # Arguments
	//  - ctx: The context to extract the user from.
	//
	// # Returns
	//  - The user wrapped in the context.
	//
	// # Errors
	//  - `ErrUnauthorized` if the user is not found in the context.
	FromContext(context.Context) (model.User, error)

	// WrapInAuthentication wraps a handler in an authentication layer.
	//
	// # Arguments
	//  - h: The handler to wrap in authentication.
	//
	// # Returns
	//  - Handler whose context contains the authenticated user.
	WrapInAuthentication(http.Handler) http.Handler
}
