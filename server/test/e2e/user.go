package e2e

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/graph/model"
)

var alice = model.User{
	ID:   uuid.NewString(),
	Name: "Alice",
}
