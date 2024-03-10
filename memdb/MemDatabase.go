package memdb

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/graph/model"
)

type MemDatabase struct {
	root  model.Folder
	users []model.User
}

func NewMemDatabase() MemDatabase {
	return MemDatabase{
		root:  model.Folder{Name: "root"},
		users: []model.User{{ID: uuid.NewString(), Name: "Peter"}},
	}
}
