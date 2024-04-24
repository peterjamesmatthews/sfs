package app

import (
	"context"
	"errors"
	"net/http"

	"pjm.dev/sfs/graph"
)

// Authenticate finds the user with the same name as the Authorization header.
//
// If no user is found, Authenticate returns graph.ErrUnauthorized.
func (a *App) Authenticate(r *http.Request) (graph.User, error) {
	return graph.User{}, errors.ErrUnsupported
}

func (a *App) WithUser(ctx context.Context, user graph.User) context.Context {
	return context.Background()
}

func (a *App) FromContext(ctx context.Context) (graph.User, error) {
	return graph.User{}, errors.ErrUnsupported
}

func (a *App) WrapInAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO authenticate
		// serve request
		h.ServeHTTP(w, r)
	})
}
