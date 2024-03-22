package mem

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/graph"
)

// Database is a simple in-memory database that implements the following interfaces:
//
//   - graph.SharedFileSystemer
//   - graph.Authenticator
//   - graph.UUIDGenerator
type Database struct {
	Root   *graph.Folder
	Users  []*graph.User
	UUIDs  []uuid.UUID
	Access []*graph.Access
}
