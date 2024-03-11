package e2e

import (
	"net/http"

	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/mem"
)

func getTestingHandler(db mem.Database) http.Handler {
	return graph.GetGQLHandler(&db, &db, &db)
}
