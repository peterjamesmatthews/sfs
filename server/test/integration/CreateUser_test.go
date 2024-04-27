package integration

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/peterjamesmatthews/crow"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name     string
		seed     crow.Database
		request  *http.Request
		response http.Response
		dump     crow.Database
	}{
		{
			name: "create first user",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h, c := New(t)

			err := c.Seed(test.seed)
			if err != nil {
				t.Fatalf("failed to seed crow: %s", err.Error())
			}

			recorder := httptest.NewRecorder()
			h.ServeHTTP(recorder, test.request)
			response := recorder.Result()

			dump, err := c.Dump()
			if err != nil {
				t.Fatalf("failed to dump crow: %s", err.Error())
			}

			if !reflect.DeepEqual(test.dump, dump) {
				t.Log("dump mismatch")
				t.Fail()
			} else if !reflect.DeepEqual(test.response, response) {
				t.Log("response mismatch")
				t.Fail()
			}
		})
	}
}
