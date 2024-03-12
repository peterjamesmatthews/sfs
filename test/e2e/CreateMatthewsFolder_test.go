package e2e

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"pjm.dev/sfs/graph/model"
	"pjm.dev/sfs/mem"
)

func TestCreateMatthewsFolder(t *testing.T) {
	uuids := []uuid.UUID{}
	for i := 0; i < 100; i++ {
		uuids = append(uuids, uuid.New())
	}

	matthew := &model.User{ID: uuids[0].String(), Name: "Matthew"}
	users := []*model.User{matthew}
	root := &model.Folder{}
	root.Children = []model.Node{}
	access := []*model.Access{
		{User: matthew, Type: model.AccessTypeRead, Target: root},
		{User: matthew, Type: model.AccessTypeWrite, Target: root},
	}

	db := &mem.Database{
		Root:   root,
		Users:  users,
		UUIDs:  uuids,
		Access: access,
	}

	request := newRequest(
		*matthew,
		"/graphql",
		GQLFileToString(t, "CreateMatthewsFolder.gql"),
	)

	handler := getTestingHandler(db)
	server := httptest.NewServer(handler)
	defer server.Close()
	request.URL.Host = server.URL

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	body := recorder.Body.String()

	// The folder should look something like this:
	strings.Contains(body, uuids[1].String())
	strings.Contains(body, "Matthew's Folder")
	strings.Contains(body, matthew.ID)
}
