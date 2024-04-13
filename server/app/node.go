package app

import (
	"github.com/google/uuid"
	"pjm.dev/sfs/db/model"
)

func (a *app) createNode(user model.User, name string, parent *model.Node) (model.Node, error) {
	var parentID *string
	if parent != nil {
		parentID = &parent.ID
	}

	node := model.Node{ID: uuid.NewString(), Name: name, Owner: user.ID, Parent: parentID}
	if err := a.db.Create(&node).Error; err != nil {
		return model.Node{}, err
	}

	return node, nil
}

func (a *app) getNodeByID(id string) (model.Node, error) {
	var node model.Node
	if err := a.db.Where("id = ?", id).First(&node).Error; err != nil {
		return model.Node{}, err
	}
	return node, nil
}
