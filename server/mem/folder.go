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

	// verify user has read access to root
	if hasAccess, err := m.has(user, model.AccessTypeRead, m.Root); err != nil {
		return model.Folder{}, fmt.Errorf("failed to check user %s has %s access on root: %w", user.ID, model.AccessTypeRead, err)
	} else if !hasAccess {
		return model.Folder{}, graph.ErrUnauthorized
	}

	return *m.Root, nil
}

func (m *Database) InsertFolder(user model.User, folder model.Folder) (model.Folder, error) {
	parent, err := m.getFolderByID(folder.Parent.ID)
	if errors.Is(err, graph.ErrNotFound) {
		return model.Folder{}, err
	} else if err != nil {
		return model.Folder{}, fmt.Errorf("failed to get parent %s: %w", folder.Parent.ID, err)
	}

	// verify user has write access to parent
	if hasAccess, err := m.has(user, model.AccessTypeWrite, m.Root); err != nil {
		return model.Folder{}, fmt.Errorf("failed to check user %s has %s access on parent %s: %w", user.ID, model.AccessTypeWrite, parent.ID, err)
	} else if !hasAccess {
		return model.Folder{}, graph.ErrUnauthorized
	}

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

	// verify user has read access to folder
	if hasAccess, err := m.has(user, model.AccessTypeRead, m.Root); err != nil {
		return model.Folder{}, fmt.Errorf("failed to check user %s has %s access on folder %s: %w", user.ID, model.AccessTypeRead, folder.ID, err)
	} else if !hasAccess {
		return model.Folder{}, graph.ErrUnauthorized
	}

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
