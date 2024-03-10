package auth

import (
	"context"
	"net/http"

	"pjm.dev/sfs/graph/model"
)

type Authenticator interface {
	Authenticate(*http.Request) (model.User, error)
	WithUser(context.Context, model.User) context.Context
	FromContext(context.Context) (model.User, error)
	WrapInAuthentication(http.Handler) http.Handler
}
