package app

import (
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
