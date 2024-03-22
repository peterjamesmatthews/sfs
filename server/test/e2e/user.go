package e2e

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/graph"
)

var alice = graph.User{
	ID:   uuid.NewString(),
	Name: "Alice",
}
