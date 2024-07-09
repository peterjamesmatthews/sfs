package app

type Auth0er interface {
	GetIDAndEmailFromToken(token string) (string, string, error)
}

func (a *App) SetAuth0er(auth0er Auth0er) {
	a.auth0 = auth0er
}
