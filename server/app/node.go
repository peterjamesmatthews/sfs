package app

import (
	"context"
	"fmt"

	"pjm.dev/sfs/db/models"
)

func (a *App) createNode(user models.User, path string) (models.Node, error) {
	segments := getPathSegments(path)
	if len(segments) == 0 {
		return models.Node{}, fmt.Errorf("path must contain at least one segment")
	}

	// get name of node from path
	name := segments[len(segments)-1]

	// get uuid of parent from path
	var parent models.Node
	var err error // declare outside of loop to avoid shadowing
	for _, segment := range segments[:len(segments)-1] {
		parent, err = a.queries.GetNodeByOwnerNameParent(
			context.Background(),
			models.GetNodeByOwnerNameParentParams{
				Owner:  user.ID,
				Name:   segment,
				Parent: parent.ID,
			},
		)
		if err != nil {
			return models.Node{}, fmt.Errorf("failed to get parent node: %w", err)
		}
	}

	// insert node
	node, err := a.queries.CreateNode(context.Background(), models.CreateNodeParams{
		Owner:  user.ID,
		Name:   name,
		Parent: parent.ID,
	})
	if err != nil {
		return models.Node{}, fmt.Errorf("failed to create node: %w", err)
	}

	return node, nil
}
