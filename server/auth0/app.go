package auth0

// App implements app.Auth0.
type App struct{}

func (a *App) GetIDAndNameFromToken(token string) (string, string, error) {
	// request user profile from auth0
	// see https://auth0.com/docs/api/authentication#get-user-info
	return "mock-id", "mock-name", nil
}
