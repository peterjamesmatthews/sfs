package app

import (
	"strings"
)

func (a *App) getTokenFromAuthorization(auth string) string {
	return strings.TrimPrefix(auth, "Bearer ")
}
