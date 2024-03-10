// MemDatabase's implementation of the graph.SharedFileSystmer interface.
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

func (m *MemDatabase) GetNodeByID(id string) (model.Node, error) {
	if id == "" {
		return m.root, nil
	}

	nodes := m.root.Children
	for _, node := range nodes {
		if node.GetID() == id {
			return node, nil
		}

		folder, ok := node.(*model.Folder)
		if !ok {
			continue
		}

		nodes = append(nodes, folder.Children...)
	}

	return nil, fmt.Errorf("folder %s not found", id)
}

func (m *MemDatabase) GetFolderByID(id string) (model.Folder, error) {
	folder, err := m.getFolderByID(id)
	if err != nil {
		return model.Folder{}, err
	}
	return *folder, nil
}

func (m *MemDatabase) getFolderByID(id string) (*model.Folder, error) {
	if id == "" {
		return m.root, nil
	}

	nodes := m.root.Children
	for _, node := range nodes {
		folder, ok := node.(*model.Folder)
		if !ok {
			continue
		}

		if folder.ID == id {
			return folder, nil
		}

		nodes = append(nodes, folder.Children...)
	}

	return &model.Folder{}, fmt.Errorf("folder %s not found", id)
}

func (m *MemDatabase) InsertFolder(folder model.Folder) (model.Folder, error) {
	parent, err := m.getFolderByID(folder.Parent.ID)
	if err != nil {
		return model.Folder{}, fmt.Errorf("failed to get parent %s: %w", folder.Parent.ID, err)
	}
	parent.Children = append(parent.Children, folder)
	return folder, nil
}
