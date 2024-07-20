package app

import (
	"context"

	"github.com/google/uuid"
	"pjm.dev/sfs/db/models"
	"pjm.dev/sfs/graph"
)

func (a *App) getUserFromGraphUser(graphUser graph.User) (models.User, error) {
	return a.queries.GetUserByEmail(context.Background(), graphUser.Email)
}

func (a *App) getGraphUser(user models.User) graph.User {
	return graph.User{ID: uuid.UUID(user.ID.Bytes).String(), Email: user.Email}
}

func (a *App) getGraphTokens(access string, refresh string) graph.Tokens {
	return graph.Tokens{Access: access, Refresh: refresh}
}

func (a *App) getGraphFolder(node models.Node) graph.Folder {
	return graph.Folder{ID: uuid.UUID(node.ID.Bytes).String(), Name: node.Name}
}
