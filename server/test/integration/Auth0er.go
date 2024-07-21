package integration

// this file provides mock's implementation of the app.Auth0er interface.

func (m *mock) GetIDAndEmailFromToken(token string) (string, string, error) {
	args := m.Called(token)
	return args.String(0), args.String(1), args.Error(2)
}
