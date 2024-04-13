package app

import (
	"errors"

	"gorm.io/gorm"
	"pjm.dev/sfs/db/model"
	"pjm.dev/sfs/graph"
)

func (a *app) toGraphUser(user model.User) graph.User {
	return graph.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (a *app) toGraphFolder(folder model.Node) graph.Folder {
	return graph.Folder{
		ID:   folder.ID,
		Name: folder.Name,
		// the following fields have their own resolvers
		// Owner
		// Parent
		// Children
	}
}

func (a *app) toGraphFile(node model.Node, file model.File) graph.File {
	return graph.File{
		ID:      file.ID,
		Name:    node.Name,
		Content: file.Content,
		// the following fields have their own resolvers
		// Owner
		// Parent
	}
}

func (a *app) toGraphNode(node model.Node) (graph.Node, error) {
	file, err := a.getFileByNode(node)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return a.toGraphFolder(node), nil
	} else if err != nil {
		return nil, err
	}
	return a.toGraphFile(node, file), nil
}
