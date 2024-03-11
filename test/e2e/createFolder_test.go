package e2e

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
	"pjm.dev/sfs/graph/model"
	"pjm.dev/sfs/memdb"
)

func TestCreateFolder(t *testing.T) {
	uuids := []uuid.UUID{}
	for i := 0; i < 100; i++ {
		uuids = append(uuids, uuid.New())
	}

	tests := []struct {
		name      string
		seed      memdb.MemDatabase
		requestor model.User
		request   *http.Request
		response  *http.Response
		want      memdb.MemDatabase
	}{
		{
			name: "empty file system",
			seed: memdb.MemDatabase{
				Root:  &root,
				Users: []*model.User{&alice},
				UUIDs: uuids,
			},
			requestor: alice,
			request: httptest.NewRequest(
				http.MethodPost,
				"/graphql",
				strings.NewReader(`{"query":"mutation{createFolder(name:"Foobar"){id}}}`),
			),
			want: memdb.MemDatabase{
				Root: &model.Folder{
					ID: root.ID,
					Children: []model.Node{&model.Folder{
						ID:       uuids[0].String(),
						Name:     "Foo",
						Owner:    &alice,
						Parent:   &root,
						Children: []model.Node{},
					}},
				},
				Users: []*model.User{&alice},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := getTestingHandler(t, test.seed)

			test.request.Header.Set("Content-Type", "application/json")

			server := httptest.NewServer(handler)
			defer server.Close()
			test.request.URL.Host = server.URL

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, test.request)
			response := recorder.Result()

			if test.response != nil && !reflect.DeepEqual(test.response, response) {
				t.Error("response mismatch")
			}

			if !reflect.DeepEqual(test.want, test.seed) {
				t.Error("database mismatch")
			}
		})
	}
}
