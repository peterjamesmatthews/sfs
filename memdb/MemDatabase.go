package memdb

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/graph/model"
)

type MemDatabase struct {
	root  *model.Folder
	users []*model.User
}

func NewDatabase() MemDatabase {
	return MemDatabase{}
}

func NewSeededDatabase() MemDatabase {
	peter := &model.User{ID: uuid.NewString(), Name: "Peter"}
	users := []*model.User{peter}

	root := &model.Folder{
		ID:    uuid.NewString(),
		Owner: peter,
	}

	foo := &model.Folder{
		ID:     uuid.NewString(),
		Name:   "Foo",
		Owner:  peter,
		Parent: root,
	}

	bar := &model.File{
		ID:      uuid.NewString(),
		Name:    "Bar",
		Owner:   peter,
		Parent:  root,
		Content: "Hello World!",
	}

	root.Children = []model.Node{foo, bar}

	return MemDatabase{root: root, users: users}
}
