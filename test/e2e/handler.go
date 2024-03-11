package e2e

import (
	"net/http"
	"testing"

	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/memdb"
)

func getTestingHandler(t *testing.T, db memdb.MemDatabase) http.Handler {
	t.Helper()
	return graph.GetGQLHandler(&db, &db, &db)
}
