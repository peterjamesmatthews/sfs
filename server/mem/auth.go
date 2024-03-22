// MemDatabase's implementation of the auth.Authenticator interface.
package mem

import (
	"context"
	"errors"
	"net/http"

	"pjm.dev/sfs/graph"
)

func (m *Database) Authenticate(r *http.Request) (graph.User, error) {
	// get name header from request
	name := r.Header.Get("Authorization")
	if name == "" {
		return graph.User{}, graph.ErrUnauthorized
	}

	// find user by name
	for _, user := range m.Users {
		if user != nil && user.Name == name {
			return *user, nil
		}
	}

	// user not found
	return graph.User{}, graph.ErrNotFound
}

type userContextKeyType string

const userContextKey userContextKeyType = "user"

func (m *Database) WithUser(ctx context.Context, user graph.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func (m *Database) FromContext(ctx context.Context) (graph.User, error) {
	user, ok := ctx.Value(userContextKey).(graph.User)
	if !ok {
		return graph.User{}, graph.ErrUnauthorized
	}

	return user, nil
}

func (m *Database) WrapInAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authenticate user
		user, err := m.Authenticate(r)
		if errors.Is(err, graph.ErrUnauthorized) {
			// serve unauthenticated request
		} else if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		} else {
			// wrap user in context
			ctx := m.WithUser(r.Context(), user)
			r = r.WithContext(ctx)
		}

		// serve request
		h.ServeHTTP(w, r)
	})
}
