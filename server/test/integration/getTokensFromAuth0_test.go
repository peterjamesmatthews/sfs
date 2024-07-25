package integration

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"pjm.dev/sfs/db/models"
)

// TestGetTokensFromAuth0 tests server's handling of the getTokensFromAuth0 query.
func TestGetTokensFromAuth0(t *testing.T) {
	t.Parallel()
	server, mock, stack := newTestServer(t)

	url, err := url.Parse(server.URL + "/graph")
	if err != nil {
		t.Fatalf("failed to parse server URL: %v", err)
	}

	mockToken := "mockToken"

	tests := []struct {
		test
		setupMock      func(m *stackMock)
		assertResponse func(t *testing.T, r *http.Response)
		assertDatabase func(t *testing.T, q *models.Queries)
	}{
		{
			setupMock: func(m *stackMock) {
				m.
					On("GetIDAndEmailFromToken", mockToken).
					Return("auth0|mock-auth0-id", "mock@pjm.dev", nil)
			},
			test: test{
				name: "creates new user",
				request: &http.Request{
					Method: http.MethodPost,
					URL:    url,
					Header: http.Header{"Content-Type": []string{"application/json"}},
					Body:   io.NopCloser(strings.NewReader(fmt.Sprintf(`{"query":"query { getTokensFromAuth0(token: \"%s\") { access refresh } }"}`, mockToken))),
				},
				response: &http.Response{StatusCode: http.StatusOK},
			},
			assertResponse: func(t *testing.T, r *http.Response) {
				assert.Equal(
					t, http.StatusOK, r.StatusCode,
					"unexpected status code",
				)

				// read response body
				bytes, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("failed to read response body: %v", err)
				}
				body := string(bytes)

				assert.NotContains(t, body, "errors", "unexpected errors in response")
				assert.Contains(t, body, "access", "access token missing from response")
				assert.Contains(t, body, "refresh", "refresh token missing from response")
			},
			assertDatabase: func(t *testing.T, q *models.Queries) {
				user, err := q.GetUserByEmail(context.Background(), "mock@pjm.dev")
				assert.NoError(t, err, "failed to get user from database: %v", err)
				// assert user has correct auth0ID and email
				assert.Equal(t, "auth0|mock-auth0-id", user.Auth0ID, "unexpected auth0ID")
				assert.Equal(t, "mock@pjm.dev", user.Email, "unexpected email")
			},
		},
		{
			test: test{
				name: "authenticates existing user",
				request: &http.Request{
					Method: http.MethodPost,
					URL:    url,
					Header: http.Header{"Content-Type": []string{"application/json"}},
					Body:   io.NopCloser(strings.NewReader(fmt.Sprintf(`{"query":"query { getTokensFromAuth0(token: \"%s\") { access refresh } }"}`, mockToken))),
				},
				response: &http.Response{StatusCode: http.StatusOK},
			},
			setupMock: func(m *stackMock) {
				m.
					On("GetIDAndEmailFromToken", mockToken).
					Return("auth0|mock-auth0-id", "mock@pjm.dev", nil)
			},
			assertResponse: func(t *testing.T, r *http.Response) {
				assert.Equal(
					t, http.StatusOK, r.StatusCode,
					"unexpected status code",
				)

				// read response body
				bytes, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("failed to read response body: %v", err)
				}
				body := string(bytes)

				assert.NotContains(t, body, "errors", "unexpected errors in response")
				assert.Contains(t, body, "access", "access token missing from response")
				assert.Contains(t, body, "refresh", "refresh token missing from response")
			},
			assertDatabase: func(t *testing.T, q *models.Queries) {
				user, err := q.GetUserByEmail(context.Background(), "mock@pjm.dev")
				assert.NoError(t, err, "failed to get user from database: %v", err)
				// assert user has correct auth0ID and email
				assert.Equal(t, "auth0|mock-auth0-id", user.Auth0ID, "unexpected auth0ID")
				assert.Equal(t, "mock@pjm.dev", user.Email, "unexpected email")
			},
		},
		{
			test: test{
				name: "bails on auth0 error",
				request: &http.Request{
					Method: http.MethodPost,
					URL:    url,
					Header: http.Header{"Content-Type": []string{"application/json"}},
					Body:   io.NopCloser(strings.NewReader(fmt.Sprintf(`{"query":"query { getTokensFromAuth0(token: \"%s\") { access refresh } }"}`, "badToken"))),
				},
				response: &http.Response{StatusCode: http.StatusOK},
			},
			setupMock: func(m *stackMock) {
				m.
					On("GetIDAndEmailFromToken", "badToken").
					Return("", "", errors.New("mock error"))
			},
			assertResponse: func(t *testing.T, r *http.Response) {
				assert.Equal(
					t, http.StatusOK, r.StatusCode,
					"unexpected status code",
				)

				// read response body
				bytes, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("failed to read response body: %v", err)
				}
				body := string(bytes)

				assert.Contains(t, body, "errors", "mock-error")
			},
			assertDatabase: func(t *testing.T, q *models.Queries) {
				n, err := q.GetCountOfUsers(context.Background())
				assert.NoError(t, err, "failed to get count of users from database")
				// only user should be from previous tests
				assert.Equal(t, int64(1), n, "unexpected number of users in database")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup mocks (if any)
			if test.setupMock != nil {
				test.setupMock(mock)
				defer func() {
					// perform assertions on mock
					mock.AssertExpectations(t)

					// reset mocks
					mock.ExpectedCalls = nil
				}()
			}

			// send request
			got, err := server.Client().Do(test.request)
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}

			if test.assertResponse != nil {
				// perform assertions on response
				test.assertResponse(t, got)
			}

			if test.assertDatabase != nil {
				// assert user is in database
				test.assertDatabase(t, stack.App.Queries)
			}
		})
	}
}
