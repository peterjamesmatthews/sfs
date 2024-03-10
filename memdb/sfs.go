// MemDatabase's implementation of the graph.SharedFileSystmer interface.
package memdb

import (
	"fmt"

	"pjm.dev/sfs/graph/model"
)

func (m *MemDatabase) GetRoot() (model.Folder, error) {
	return m.root, nil
}

func (m *MemDatabase) GetNodeByID(id string) (model.Node, error) {
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
	nodes := m.root.Children

	for _, node := range nodes {
		folder, ok := node.(model.Folder)
		if !ok {
			continue
		}

		if folder.ID == id {
			return folder, nil
		}

		nodes = append(nodes, folder.Children...)
	}

	return model.Folder{}, fmt.Errorf("folder %s not found", id)
}

func (m *MemDatabase) InsertFolder(folder model.Folder) (model.Folder, error) {
	parent := m.root

	if folder.Parent != nil {
		var err error
		parent, err = m.GetFolderByID(folder.Parent.ID)
		if err != nil {
			return model.Folder{}, fmt.Errorf("failed to get parent %s: %w", folder.Parent.ID, err)
		}
	}

	parent.Children = append(parent.Children, folder)

	return folder, nil
}
