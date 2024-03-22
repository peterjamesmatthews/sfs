package mem

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/graph"
)

type Database struct {
	Root   *graph.Folder
	Users  []*graph.User
	UUIDs  []uuid.UUID
	Access []*graph.Access
}
