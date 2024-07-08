package app

type Auth0er interface {
	GetIDAndEmailFromToken(token string) (string, string, error)
}
