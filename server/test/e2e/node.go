package e2e

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"pjm.dev/sfs/mem"
)

func TestGetNodeByURI(t *testing.T) {
	tests := []struct {
		name     string
		seed     mem.Database
		request  *http.Request
		response *http.Response
		want     mem.Database
	}{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := test.seed
			handler := getTestingHandler(&db)

			server := httptest.NewServer(handler)
			defer server.Close()
			test.request.URL.Host = server.URL

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, test.request)
			response := recorder.Result()

			if test.response != nil && !reflect.DeepEqual(test.response, response) {
				t.Error("response mismatch")
			}

			if !reflect.DeepEqual(test.want.Root, test.seed.Root) {
				t.Error("database mismatch")
			}
		})
	}
}
