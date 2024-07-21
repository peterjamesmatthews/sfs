package integration

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"pjm.dev/sfs/db/models"
	"pjm.dev/sfs/graph"
)

func TestGetNodeFromPath(t *testing.T) {
	t.Parallel()
	server, mock, stack := newTestServer(t)
	defer server.Close()

	// seed database with user
	user, err := stack.App.Queries.CreateUser(
		context.Background(),
		models.CreateUserParams{
			Email:   "mock-user@pjm.dev",
			Auth0ID: pgtype.Text{String: "auth0|mock-user-auth0-id", Valid: true},
		},
	)
	if err != nil {
		t.Fatalf("unable to seed user: %v", err)
	}

	folder, err := stack.App.CreateFolder(graph.User{Email: user.Email}, "/foo")
	if err != nil {
		t.Fatalf("unable to seed folder: %v", err)
	}

	// get tokens for user
	accessToken, _, err := stack.App.GetTokensForUser(user)
	if err != nil {
		t.Fatalf("unable to get tokens for mock user: %v", err)
	}

	url, err := url.Parse(server.URL + "/graph")
	if err != nil {
		t.Fatalf("failed to parse server URL: %v", err)
	}

	tests := []test{
		{
			name: "root",
			request: &http.Request{
				Method: http.MethodPost,
				URL:    url,
				Header: http.Header{
					"Authorization": []string{"Bearer " + accessToken},
					"Content-Type":  []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`{ "query": "query { getNodeFromPath(path: \"/\") { name } }" }`)),
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{ "data": { "getNodeFromPath": { "name": "" } } }`)),
			},
		},
		{
			name: "not found",
			request: &http.Request{
				Method: http.MethodPost,
				URL:    url,
				Header: http.Header{
					"Authorization": []string{"Bearer " + accessToken},
					"Content-Type":  []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`{ "query": "query { getNodeFromPath(path: \"/does/not/exist\") { name } }" }`)),
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{ "data": null, "errors": [{ "message": "node not found", "path": ["getNodeFromPath"] }] }`)),
			},
		},
		{
			name: "folder",
			request: &http.Request{
				Method: http.MethodPost,
				URL:    url,
				Header: http.Header{
					"Authorization": []string{"Bearer " + accessToken},
					"Content-Type":  []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`{ "query": "query { getNodeFromPath(path: \"/foo\") { name } }" }`)),
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(fmt.Sprintf(`{ "data": { "getNodeFromPath": { "name": "%s" } } }`, folder.Name))),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := server.Client().Do(test.request)
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}

			// perform assertions on response status code
			assert.Equal(t,
				test.response.StatusCode, got.StatusCode,
				"unexpected response status code",
			)

			bytes, err := io.ReadAll(got.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}
			gotBody := string(bytes)

			bytes, err = io.ReadAll(test.response.Body)
			if err != nil {
				t.Fatalf("failed to read expected response body: %v", err)
			}
			wantBody := string(bytes)

			assert.JSONEq(t,
				wantBody, gotBody,
				"unexpected response body",
			)

			// perform assertions on mock
			mock.AssertExpectations(t)
		})
	}
}
