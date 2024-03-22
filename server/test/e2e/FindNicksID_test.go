package e2e

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/mem"
)

func TestFindNicksID(t *testing.T) {
	nick := &graph.User{ID: uuid.NewString(), Name: "Nick"}
	users := []*graph.User{nick}
	root := &graph.Folder{}
	root.Children = []graph.Node{
		&graph.File{
			ID:      uuid.NewString(),
			Name:    "Nick's File",
			Owner:   nick,
			Parent:  root,
			Content: "Hello World!",
		},
	}

	access := []*graph.Access{
		{User: nick, Type: graph.AccessTypeRead, Target: root},
	}

	db := mem.Database{
		Root:   root,
		Users:  users,
		Access: access,
	}

	request := newRequest(
		*nick,
		"/graphql",
		GQLFileToString(t, "FindNicksID.gql"),
	)

	handler := getTestingHandler(&db)
	server := httptest.NewServer(handler)
	defer server.Close()
	request.URL.Host = server.URL

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	body := recorder.Body.String()

	// search body for nick's id
	if !strings.Contains(body, nick.ID) {
		t.Errorf("Response body does not contain Nick's ID %s", nick.ID)
	}
}
