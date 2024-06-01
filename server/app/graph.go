package app

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/db/models"
	"pjm.dev/sfs/graph"
)

func (a *App) getGraphUser(user models.User) graph.User {
	return graph.User{ID: uuid.UUID(user.ID.Bytes).String(), Name: user.Name}
}

func (a *App) getGraphTokens(access string, refresh string) graph.Tokens {
	return graph.Tokens{Access: access, Refresh: refresh}
}
