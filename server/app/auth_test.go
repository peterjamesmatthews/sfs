package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetIDAndNameFromToken(t *testing.T) {
	// Create a new instance of the App
	app := &App{}

	// Create a test server with an insecure TLS configuration
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Write the response body
		response := map[string]interface{}{
			"sub":  "123",
			"name": "John Doe",
		}
		json.NewEncoder(w).Encode(response)
	}))
	client := server.Client()
	http.DefaultClient = client
	defer server.Close()

	// Set the test server URL in the App
	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse server URL: %v", err)
	}

	// Set the AUTH0_DOMAIN in the App
	app.config.AUTH0_DOMAIN = serverURL.Host

	// Call the getIDAndNameFromToken function with a valid token
	id, name, err := app.getIDAndNameFromToken("valid_token")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify the returned ID and name
	expectedID := "123"
	expectedName := "John Doe"
	if id != expectedID {
		t.Errorf("Expected ID to be %s, but got %s", expectedID, id)
	}
	if name != expectedName {
		t.Errorf("Expected name to be %s, but got %s", expectedName, name)
	}
}
