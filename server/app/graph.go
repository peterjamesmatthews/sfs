package app

import (
	"context"
	"fmt"

	"pjm.dev/sfs/db/models"
	"pjm.dev/sfs/graph"
)

func (a *App) getUserFromGraphUser(graphUser graph.User) (models.User, error) {
	return a.Queries.GetUserByEmail(context.Background(), graphUser.Email)
}

func (a *App) getGraphUser(user models.User) graph.User {
	return graph.User{ID: user.ID.String(), Email: user.Email}
}

func (a *App) getGraphTokens(access string, refresh string) graph.Tokens {
	return graph.Tokens{Access: access, Refresh: refresh}
}

func (a *App) getGraphFolder(node models.Node) graph.Folder {
	return graph.Folder{ID: node.ID.String(), Name: node.Name}
}

func (a *App) getGraphFile(node models.Node, file models.File) graph.File {
	return graph.File{
		ID:      node.ID.String(),
		Name:    node.Name,
		Content: string(file.Content),
	}
}

func (a *App) getGraphNode(node models.Node) (graph.Node, error) {
	file, err := a.Queries.GetFileByNode(context.Background(), node.ID)
	if a.isNotFoundError(err) {
		return a.getGraphFolder(node), nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get node: %w", err)
	}
	return a.getGraphFile(node, file), nil
}
