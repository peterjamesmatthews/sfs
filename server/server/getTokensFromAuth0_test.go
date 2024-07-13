package server_test

import (
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
	server, mock, db := newTestServer(t)
	defer server.Close()

	token := "mock-new-user-token"
	id := "mock-new-user-id"
	email := "mock-new-user-email"

	mock.
		On("GetIDAndEmailFromToken", token).
		Return(id, email, nil)

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

		// perform assertions on database
		dump := dumpDatabase(t, db)
		if err != nil {
			t.Fatalf("failed to dump database: %v", err)
		}
		assert.Contains(t, dump, id, "user ID missing from database dump: %s", dump)
		assert.Contains(t, dump, email, "user email missing from database dump: %s", dump)
	}
}
