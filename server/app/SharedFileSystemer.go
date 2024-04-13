package app

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
	"pjm.dev/sfs/db/model"
	"pjm.dev/sfs/graph"
)

func (a *app) CreateUser(name string) (graph.User, error) {
	// validate name
	if name == "" {
		return graph.User{}, errors.New("name cannot be empty")
	}

	// insert user
	user := model.User{Name: name}
	if err := a.db.Create(&user).Error; errors.Is(err, gorm.ErrDuplicatedKey) {
		return graph.User{}, graph.ErrConflict
	} else if err != nil {
		log.Printf("failed to create user: %v", err)
		return graph.User{}, fmt.Errorf("failed to create user")
	}

	// return graph.User
	return a.toGraphUser(user), nil
}

func (a *app) GetNodeByURI(user graph.User, uri string) (graph.Node, error) {
	return nil, errors.New("not implemented")
}

func (a *app) RenameNode(user graph.User, id string, name string) (graph.Node, error) {
	return nil, errors.New("not implemented")
}

func (a *app) MoveNode(user graph.User, id string, dstID string) (graph.Node, error) {
	return nil, errors.New("not implemented")
}

func (a *app) GetRoot(user graph.User) (graph.Folder, error) {
	return graph.Folder{}, errors.New("not implemented")
}

func (a *app) InsertFolder(user graph.User, folder graph.Folder) (graph.Folder, error) {
	return graph.Folder{}, errors.New("not implemented")
}

func (a *app) GetFolderByID(user graph.User, id string) (graph.Folder, error) {
	return graph.Folder{}, errors.New("not implemented")
}

func (a *app) InsertFile(user graph.User, file graph.File) (graph.File, error) {
	return graph.File{}, errors.New("not implemented")
}

func (a *app) WriteFile(user graph.User, fileID string, content string) (graph.File, error) {
	return graph.File{}, errors.New("not implemented")
}

func (a *app) GetUserByID(id string) (graph.User, error) {
	return graph.User{}, errors.New("not implemented")
}
