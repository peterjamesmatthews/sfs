package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"pjm.dev/sfs/db/models"
)

func (a *App) createNode(user models.User, path string) (models.Node, error) {
	// get segments of path
	segments := getPathSegments(path)
	if len(segments) == 0 {
		return models.Node{}, fmt.Errorf("path must contain at least one segment")
	}

	// get name of node from path
	name := segments[len(segments)-1]

	// get uuid of parent from path
	var parent models.Node
	var err error // declare outside of loop to avoid shadowing parent
	for _, segment := range segments[:len(segments)-1] {
		parent, err = a.Queries.GetNodeByOwnerNameParent(
			context.Background(),
			models.GetNodeByOwnerNameParentParams{
				Owner:  user.ID,
				Name:   segment,
				Parent: uuid.NullUUID{UUID: parent.ID, Valid: parent.ID != uuid.Nil},
			},
		)
		if err != nil {
			return models.Node{}, fmt.Errorf("failed to get parent node: %w", err)
		}
	}

	// insert node
	node, err := a.Queries.CreateNode(context.Background(), models.CreateNodeParams{
		Owner:  user.ID,
		Name:   name,
		Parent: uuid.NullUUID{UUID: parent.ID, Valid: parent.ID != uuid.Nil},
	})
	if err != nil {
		return models.Node{}, fmt.Errorf("failed to create node: %w", err)
	}

	return node, nil
}

func (a *App) getNodeFromPath(user models.User, path string) (models.Node, error) {
	// get segments of path
	segments := getPathSegments(path)
	if len(segments) == 0 {
		return models.Node{}, fmt.Errorf("path must contain at least one segment")
	}

	// for each segment, get node by owner, name, and parentID
	var parentID uuid.UUID
	var node models.Node
	var err error // declare outside of loop to avoid shadowing parent
	for _, segment := range segments {
		node, err = a.Queries.GetNodeByOwnerNameParent(
			context.Background(),
			models.GetNodeByOwnerNameParentParams{
				Owner:  user.ID,
				Name:   segment,
				Parent: uuid.NullUUID{UUID: parentID, Valid: parentID != uuid.Nil},
			},
		)
		if err != nil {
			return models.Node{}, fmt.Errorf("failed to get node: %w", err)
		}
		parentID = node.ID
	}

	return node, nil
}
