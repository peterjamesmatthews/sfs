package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIDAndNameFromToken(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		response auth0UserInfoResponse
	}{
		{
			name:     "valid token",
			token:    "mock-token",
			response: auth0UserInfoResponse{Sub: "foo", Email: "bar@example.com"},
		},
	}

	// Create a new instance of the App
	app := &App{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a test server with an insecure TLS configuration
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// assert this is a GET request
				assert.Equal(t, r.Method, http.MethodGet, "unexpected method")

				// assert this request's path is /userinfo
				assert.Equal(t, r.URL.Path, "/userinfo", "unexpected path")

				// assert the Authorization header is of the form "Bearer <test.token>"
				assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+test.token, "unexpected Authorization header")

				// Set the response headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Write the response body
				response := auth0UserInfoResponse{
					Sub:   test.response.Sub,
					Email: test.response.Email,
				}
				json.NewEncoder(w).Encode(response)
			}))
			defer server.Close()

			// Set the test server's client as the default client
			client := server.Client()
			http.DefaultClient = client

			// Set the test server's URL as the AUTH0_DOMAIN
			serverURL, err := url.Parse(server.URL)
			if err != nil {
				t.Fatalf("Failed to parse server URL: %v", err)
			}
			app.config.AUTH0_DOMAIN = serverURL.Host

			// Call the getIDAndNameFromToken function with a valid token
			id, name, err := app.getIDAndEmailFromToken(test.token)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Verify the returned ID and name
			assert.Equal(t, test.response.Sub, id, "unexpected ID")
			assert.Equal(t, test.response.Email, name, "unexpected name")
		})
	}
}
