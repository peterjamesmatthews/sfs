package app

import (
	"crypto/sha256"
	"strings"
)

func (a *App) getTokenFromAuthorization(auth string) string {
	return strings.TrimPrefix(auth, "Bearer ")
}

func (a *App) hashToken(token string) []byte {
	arr := sha256.Sum256([]byte(token))
	return arr[:]
}
