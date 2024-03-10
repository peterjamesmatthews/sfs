// MemDatabase's implementation of the auth.Authenticator interface.
package memdb

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"pjm.dev/sfs/graph/model"
)

var errMissingAuthorizationHeader = errors.New("missing Authorization header")

func (m *MemDatabase) Authenticate(r *http.Request) (model.User, error) {
	// get name header from request
	name := r.Header.Get("Authorization")
	if name == "" {
		return model.User{}, errMissingAuthorizationHeader
	}

	// find user by name
	for _, user := range m.users {
		if user.Name == name {
			return user, nil
		}
	}

	// user not found
	return model.User{}, fmt.Errorf("user %s not found", name)
}

type userContextKeyType string

const userContextKey userContextKeyType = "user"

func (m *MemDatabase) WithUser(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func (m *MemDatabase) FromContext(ctx context.Context) (model.User, error) {
	user, ok := ctx.Value(userContextKey).(model.User)
	if !ok {
		return model.User{}, errors.New("user not found in context")
	}

	return user, nil
}

func (m *MemDatabase) WrapInAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authenticate user
		user, err := m.Authenticate(r)
		if errors.Is(err, errMissingAuthorizationHeader) {
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
