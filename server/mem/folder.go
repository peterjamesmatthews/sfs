package mem

import (
	"errors"
	"fmt"

	"pjm.dev/sfs/graph"
)

func (m *Database) GetRoot(user graph.User) (graph.Folder, error) {
	if m.Root == nil {
		return graph.Folder{}, graph.ErrNotFound
	}

	// verify user has read access to root
	if hasAccess, err := m.has(user, graph.AccessTypeRead, m.Root); err != nil {
		return graph.Folder{}, fmt.Errorf("failed to check user %s has %s access on root: %w", user.ID, graph.AccessTypeRead, err)
	} else if !hasAccess {
		return graph.Folder{}, graph.ErrUnauthorized
	}

	return *m.Root, nil
}

func (m *Database) InsertFolder(user graph.User, folder graph.Folder) (graph.Folder, error) {
	parent, err := m.getFolderByID(folder.Parent.ID)
	if errors.Is(err, graph.ErrNotFound) {
		return graph.Folder{}, err
	} else if err != nil {
		return graph.Folder{}, fmt.Errorf("failed to get parent %s: %w", folder.Parent.ID, err)
	}

	// verify user has write access to parent
	if hasAccess, err := m.has(user, graph.AccessTypeWrite, m.Root); err != nil {
		return graph.Folder{}, fmt.Errorf("failed to check user %s has %s access on parent %s: %w", user.ID, graph.AccessTypeWrite, parent.ID, err)
	} else if !hasAccess {
		return graph.Folder{}, graph.ErrUnauthorized
	}

	folder.Parent = parent
	parent.Children = append(parent.Children, &folder)

	return folder, nil
}

func (m *Database) GetFolderByID(user graph.User, id string) (graph.Folder, error) {
	folder, err := m.getFolderByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return graph.Folder{}, err
	} else if err != nil {
		return graph.Folder{}, fmt.Errorf("failed to get folder %s: %w", folder.Parent.ID, err)
	}

	// verify user has read access to folder
	if hasAccess, err := m.has(user, graph.AccessTypeRead, m.Root); err != nil {
		return graph.Folder{}, fmt.Errorf("failed to check user %s has %s access on folder %s: %w", user.ID, graph.AccessTypeRead, folder.ID, err)
	} else if !hasAccess {
		return graph.Folder{}, graph.ErrUnauthorized
	}

	return *folder, nil
}

func (m *Database) getFolderByID(id string) (*graph.Folder, error) {
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get node %s: %w", id, err)
	}

	folder, ok := node.(*graph.Folder)
	if !ok {
		return nil, fmt.Errorf("node %s is not a folder", id)
	}

	return folder, nil
}
