package app

import "pjm.dev/sfs/db/model"

func (a *App) getUserByID(id string) (model.User, error) {
	user := model.User{ID: id}
	err := a.db.Where(&user).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
