package app

import (
	"context"
	"errors"
	"net/http"

	"pjm.dev/sfs/db/model"
	"pjm.dev/sfs/graph"
)

// Authenticate finds the user with the same name as the Authorization header.
//
// If no user is found, Authenticate returns graph.ErrUnauthorized.
func (a *app) Authenticate(r *http.Request) (graph.User, error) {
	name := r.Header.Get("Authorization")
	if name == "" {
		return graph.User{}, graph.ErrUnauthorized
	}

	var user model.User
	if err := a.db.Where("name = ?", name).First(&user).Error; err != nil {
		return graph.User{}, graph.ErrUnauthorized
	}

	return a.toGraphUser(user), nil
}

type appContextKeyType string

const userContextKey appContextKeyType = "user"

func (a *app) WithUser(ctx context.Context, user graph.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func (a *app) FromContext(ctx context.Context) (graph.User, error) {
	user, ok := ctx.Value(userContextKey).(graph.User)
	if !ok {
		return graph.User{}, graph.ErrUnauthorized
	}

	return user, nil
}

func (a *app) WrapInAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authenticate user
		user, err := a.Authenticate(r)
		if errors.Is(err, graph.ErrUnauthorized) {
			// serve unauthenticated request
		} else if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		} else {
			// wrap user in context
			ctx := a.WithUser(r.Context(), user)
			r = r.WithContext(ctx)
		}

		// serve request
		h.ServeHTTP(w, r)
	})
}
