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
