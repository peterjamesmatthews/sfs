package app

import (
	"errors"

	"gorm.io/gorm"
	"pjm.dev/sfs/db/model"
)

type accessType string

const (
	read  accessType = "read"
	write accessType = "write"
)

func (a *app) giveAccess(user model.User, read accessType, folder model.Node) error {
	access := model.Access{User: user.ID, Type: string(read), Node: folder.ID}
	err := a.db.Where(&access).First(&access).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return a.db.Create(&access).Error
	}

	return err
}

func (a *app) hasAccess(user model.User, at accessType, node model.Node) (bool, error) {
	access := model.Access{User: user.ID, Type: string(at), Node: node.ID}
	err := a.db.Where(&access).First(&access).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
