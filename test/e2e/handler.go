package e2e

import (
	"net/http"
	"testing"

	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/mem"
)

func getTestingHandler(t *testing.T, db mem.Database) http.Handler {
	t.Helper()
	return graph.GetGQLHandler(&db, &db, &db)
}
