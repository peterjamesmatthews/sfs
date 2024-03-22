package e2e

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/mem"
)

func TestCreateFolder(t *testing.T) {
	uuids := []uuid.UUID{}
	for i := 0; i < 100; i++ {
		uuids = append(uuids, uuid.New())
	}

	tests := []struct {
		name     string
		seed     mem.Database
		request  *http.Request
		response *http.Response
		want     mem.Database
	}{
		{
			name: "empty file system",
			seed: mem.Database{
				Root:  &root,
				Users: []*graph.User{&alice},
				UUIDs: uuids,
				Access: []*graph.Access{
					{User: &alice, Type: graph.AccessTypeRead, Target: &root},
					{User: &alice, Type: graph.AccessTypeWrite, Target: &root},
				},
			},
			request: newRequest(
				alice,
				"/graphql",
				`{"query":"mutation CreateFolder { createFolder(name:\"Foobar\") { id } }","operationName":"CreateFolder"}`,
			),
			want: mem.Database{
				Root: &graph.Folder{
					ID: root.ID,
					Children: []graph.Node{
						&graph.Folder{
							ID:     uuids[0].String(),
							Name:   "Foobar",
							Owner:  &alice,
							Parent: &root,
						},
					},
				},
				Users: []*graph.User{&alice},
				Access: []*graph.Access{
					{User: &alice, Type: graph.AccessTypeRead, Target: &root},
					{User: &alice, Type: graph.AccessTypeWrite, Target: &root},
				},
			},
		},
	}

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
