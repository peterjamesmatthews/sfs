package server_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test represents a single test for server_test tests.
type test struct {
	// identifier for the test
	name string

	// database state before the test
	seed *os.File

	// request to send to the server
	request *http.Request

	// response to expect from the server
	response *http.Response

	// database state after the test
	dump *os.File
}

// TestGetTokensFromAuth0 tests server's handling of the getTokensFromAuth0 query.
func TestGetTokensFromAuth0(t *testing.T) { // TODO parametrize test with more cases
	server, mock := newTestServer(t)
	defer server.Close()

	// test a user is created when logging in for the first time
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

	request := &http.Request{
		Method: http.MethodPost,
		URL:    url,
		Header: http.Header{"Content-Type": []string{"application/json"}},
		Body:   io.NopCloser(strings.NewReader(fmt.Sprintf(`{"query":"query { getTokensFromAuth0(token: \"%s\") { access refresh } }"}`, token))), // TODO ref to test fixture
	}
	want := &http.Response{StatusCode: http.StatusOK}

	got, err := server.Client().Do(request)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}

	assert.Equal(t, want.StatusCode, got.StatusCode, "unexpected status code: want %d, got %d", want.StatusCode, got.StatusCode)

	bytes, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	body := string(bytes)

	assert.NotContains(t, body, "errors", "unexpected errors in response: %s", body)
	assert.Contains(t, body, "access", "access token missing from response: %s", body)
	assert.Contains(t, body, "refresh", "refresh token missing from response: %s", body)

	mock.AssertExpectations(t)

	// TODO assert new user was created in database
}

func (m *mock) GetIDAndEmailFromToken(token string) (string, string, error) {
	args := m.Called(token)
	return args.String(0), args.String(1), args.Error(2)
}
