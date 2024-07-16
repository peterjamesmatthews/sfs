package server_test

import (
	"io"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNodeFromPath(t *testing.T) {
	server, mock, db := newTestServer(t)
	defer server.Close()

	_, err := url.Parse(server.URL + "/graph")
	if err != nil {
		t.Fatalf("failed to parse server URL: %v", err)
	}

	tests := []test{}
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

			// perform assertions on database
			gotDump := dumpDatabase(t, db)
			bytes, err = io.ReadAll(test.dump)
			if err != nil {
				t.Fatalf("failed to read expected database dump: %v", err)
			}
			wantDump := string(bytes)
			assert.Equal(t,
				wantDump, gotDump,
				"unexpected database state",
			)
		})
	}
}
