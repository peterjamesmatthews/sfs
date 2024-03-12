package mem

import (
	"errors"
	"fmt"

	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/graph/model"
)

func (m *Database) GetRoot(user model.User) (model.Folder, error) {
	if m.Root == nil {
		return model.Folder{}, graph.ErrNotFound
	}

	// TODO verify user has read access to root

	return *m.Root, nil
}

func (m *Database) InsertFolder(user model.User, folder model.Folder) (model.Folder, error) {
	parent, err := m.getFolderByID(folder.Parent.ID)
	if errors.Is(err, graph.ErrNotFound) {
		return model.Folder{}, err
	} else if err != nil {
		return model.Folder{}, fmt.Errorf("failed to get parent %s: %w", folder.Parent.ID, err)
	}

	// TODO verify user has write access to parent

	folder.Parent = parent
	parent.Children = append(parent.Children, &folder)

	return folder, nil
}

func (m *Database) GetFolderByID(user model.User, id string) (model.Folder, error) {
	folder, err := m.getFolderByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return model.Folder{}, err
	} else if err != nil {
		return model.Folder{}, fmt.Errorf("failed to get folder %s: %w", folder.Parent.ID, err)
	}

	// TODO verify user has read access to folder

	return *folder, nil
}

func (m *Database) getFolderByID(id string) (*model.Folder, error) {
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get node %s: %w", id, err)
	}

	folder, ok := node.(*model.Folder)
	if !ok {
		return nil, fmt.Errorf("node %s is not a folder", id)
	}

	return folder, nil
}
