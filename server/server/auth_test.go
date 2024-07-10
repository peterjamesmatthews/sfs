package server_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func (m *mock) GetIDAndEmailFromToken(token string) (string, string, error) {
	args := m.Called(token)
	return args.String(0), args.String(1), args.Error(2)
}

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

		assert.Equal(
			t, test.response.StatusCode, got.StatusCode,
			"unexpected status code: want %d, got %d", test.response.StatusCode, got.StatusCode,
		)

		bytes, err := io.ReadAll(got.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		body := string(bytes)

		assert.NotContains(t, body, "errors", "unexpected errors in response: %s", body)
		assert.Contains(t, body, "access", "access token missing from response: %s", body)
		assert.Contains(t, body, "refresh", "refresh token missing from response: %s", body)
		mock.AssertExpectations(t)

		dump := dumpDatabase(db, t)
		if err != nil {
			t.Fatalf("failed to dump database: %v", err)
		}

		assert.Contains(t, dump, id, "user ID missing from database dump: %s", dump)
		assert.Contains(t, dump, email, "user email missing from database dump: %s", dump)
	}
}

func dumpDatabase(db *pgx.Conn, t *testing.T) string {
	cmd := exec.Command(
		"pg_dump",
		"-d", db.Config().Database,
		"-h", db.Config().Host,
		"-p", fmt.Sprintf("%d", db.Config().Port),
		"-U", db.Config().User,
	)
	cmd.Env = append(cmd.Environ(), "PGPASSWORD="+db.Config().Password)
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to dump database: %v", err)
	}
	return string(output)
}
