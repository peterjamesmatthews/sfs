package app

import (
	"fmt"

	"pjm.dev/sfs/db/model"
)

func (a *App) getFileByNode(node model.Node) (model.File, error) {
	var file model.File
	if err := a.db.Where("node = ?", node.ID).First(&file).Error; err != nil {
		return model.File{}, fmt.Errorf("failed to get file: %w", err)
	}
	return file, nil
}
