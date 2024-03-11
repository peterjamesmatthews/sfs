package memdb

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/graph/model"
)

type MemDatabase struct {
	Root  *model.Folder
	Users []*model.User
	UUIDs []uuid.UUID
}
