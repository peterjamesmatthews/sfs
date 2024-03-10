package memdb

import (
	"errors"
	"fmt"

	"pjm.dev/sfs/graph/model"
)

func (m *MemDatabase) GetRoot() (model.Folder, error) {
	if m.root == nil {
		return model.Folder{}, errors.New("nil root")
	}

	return *m.root, nil
}

func (m *MemDatabase) GetFolderByID(id string) (model.Folder, error) {
	folder, err := m.getFolderByID(id)
	if err != nil {
		return model.Folder{}, fmt.Errorf("failed to get folder %s: %w", id, err)
	}

	return *folder, nil
}

func (m *MemDatabase) getFolderByID(id string) (*model.Folder, error) {
	node, err := m.getNodeByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get node %s: %w", id, err)
	}

	folder, ok := node.(*model.Folder)
	if !ok {
		return nil, fmt.Errorf("node %s is not a folder", id)
	}

	return folder, nil
}

func (m *MemDatabase) InsertFolder(folder model.Folder) (model.Folder, error) {
	parent, err := m.getFolderByID(folder.Parent.ID)
	if err != nil {
		return model.Folder{}, fmt.Errorf("failed to get parent %s: %w", folder.Parent.ID, err)
	}

	parent.Children = append(parent.Children, &folder)

	return folder, nil
}
