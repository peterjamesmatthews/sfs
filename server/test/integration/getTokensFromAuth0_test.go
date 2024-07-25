package integration

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetTokensFromAuth0 tests server's handling of the getTokensFromAuth0 query.
func TestGetTokensFromAuth0(t *testing.T) {
	t.Parallel()
	server, mock, stack := newTestServer(t)
	defer server.Close()

	token := "mock-new-user-token"
	auth0ID := "mock-new-user-auth0-id"
	email := "mock-new-user-email"

	mock.
		On("GetIDAndEmailFromToken", token).
		Return(auth0ID, email, nil)

	url, err := url.Parse(server.URL + "/graph")
	if err != nil {
		t.Fatalf("failed to parse server URL: %v", err)
	}

	tests := []test{
		{
			name: "creates new user",
			request: &http.Request{
				Method: http.MethodPost,
				URL:    url,
				Header: http.Header{"Content-Type": []string{"application/json"}},
				Body:   io.NopCloser(strings.NewReader(fmt.Sprintf(`{"query":"query { getTokensFromAuth0(token: \"%s\") { access refresh } }"}`, token))),
			},
			response: &http.Response{StatusCode: http.StatusOK},
		},
		{
			name: "authenticates existing user",
			request: &http.Request{
				Method: http.MethodPost,
				URL:    url,
				Header: http.Header{"Content-Type": []string{"application/json"}},
				Body:   io.NopCloser(strings.NewReader(fmt.Sprintf(`{"query":"query { getTokensFromAuth0(token: \"%s\") { access refresh } }"}`, token))),
			},
			response: &http.Response{StatusCode: http.StatusOK},
		},
	}

	for _, test := range tests {
		got, err := server.Client().Do(test.request)
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}

		// perform assertions on response status code
		assert.Equal(
			t, test.response.StatusCode, got.StatusCode,
			"unexpected status code: want %d, got %d", test.response.StatusCode, got.StatusCode,
		)

		bytes, err := io.ReadAll(got.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		body := string(bytes)

		// perform assertions on response body
		assert.NotContains(t, body, "errors", "unexpected errors in response: %s", body)
		assert.Contains(t, body, "access", "access token missing from response: %s", body)
		assert.Contains(t, body, "refresh", "refresh token missing from response: %s", body)

		// perform assertions on mock
		mock.AssertExpectations(t)

		// assert user is in database
		user, err := stack.App.Queries.GetUserByEmail(context.Background(), email)
		assert.NoError(t, err, "failed to get user from database: %v", err)
		// assert user has correct auth0ID and email
		assert.Equal(t, auth0ID, user.Auth0ID, "unexpected auth0ID")
		assert.Equal(t, email, user.Email, "unexpected email")
	}
}
