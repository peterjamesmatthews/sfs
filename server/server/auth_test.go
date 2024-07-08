package server_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test represents a single test for server_test tests.
type test struct {
	// identifier for the test
	name string

	// database state before the test
	seed any

	// request to send to the server
	request *http.Request

	// response to expect from the server
	response *http.Response
}

// TestGetTokensFromAuth0 tests server's handling of the getTokensFromAuth0 query.
func TestGetTokensFromAuth0(t *testing.T) {
	server := newTestServer(t)
	defer server.Close()

	tests := []test{
		// TODO add tests
		// {
		// 	name:    "creates a new user",
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()

			response, err := server.Client().Do(test.request)
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}

			assert.Equal(t, test.response.StatusCode, response.StatusCode)
		})
	}
}
