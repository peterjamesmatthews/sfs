package mem

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/graph/model"
)

type Database struct {
	Root   *model.Folder
	Users  []*model.User
	UUIDs  []uuid.UUID
	Access []*model.Access
}
