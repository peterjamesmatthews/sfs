package app

type Auth0 interface {
	// GetIDAndNameFromToken gets the ID and name of a user from their opaque token.
	GetIDAndNameFromToken(token string) (id string, name string, err error)
}
