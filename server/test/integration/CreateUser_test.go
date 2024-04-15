package integration

import (
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name     string         // name of test
		seed     map[any][]any  // database seed
		request  *http.Request  // request to send
		response *http.Response // response to receive
		dump     map[any][]any  // database dump after handling
	}{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}
}
