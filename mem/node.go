package mem

import (
	"errors"
	"fmt"
	"slices"

	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/graph/model"
)

func (m *Database) GetNodeByID(user model.User, id string) (model.Node, error) {
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get node %s: %w", id, err)
	}

	// TODO verify user has read access to node

	return node, nil
}

func (m *Database) getNodeByID(id string) (model.Node, error) {
	if id == m.Root.ID {
		return m.Root, nil
	}

	nodes := m.Root.Children
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

	return nil, graph.ErrNotFound
}

func (m *Database) getNodeByName(name string) (model.Node, error) {
	if m.Root.Name == name {
		return m.Root, nil
	}

	nodes := m.Root.Children
	for _, node := range nodes {
		if node.GetName() == name {
			return node, nil
		}

		folder, ok := node.(*model.Folder)
		if !ok {
			continue
		}

		nodes = append(nodes, folder.Children...)
	}

	return nil, graph.ErrNotFound
}

func (m *Database) RenameNode(user model.User, id string, name string) (model.Node, error) {
	// get node
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get node %s: %w", id, err)
	}

	// TODO verify user has write access to node

	// set update node's name
	switch n := node.(type) {
	case *model.Folder:
		n.Name = name
		node = *n
	case *model.File:
		n.Name = name
		node = *n
	}

	// return renamed node
	return node, nil
}

func (m *Database) MoveNode(user model.User, id string, dstID string) (model.Node, error) {
	// get node
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to rename node %s: %w", node.GetID(), err)
	}

	// TODO verify user owns the node

	// get destination parent folder
	dst, err := m.getFolderByID(dstID)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to rename node %s: %w", node.GetID(), err)
	}

	// TODO verify user has write access to destination parent

	// get source parent folder
	src := node.GetParent()

	// remove node from src's children
	src.Children = slices.DeleteFunc(src.Children, func(child model.Node) bool {
		return child.GetID() == node.GetID()
	})

	// set node's parent to dst
	switch n := node.(type) {
	case *model.Folder:
		n.Parent = dst
		node = *n
	case *model.File:
		n.Parent = dst
		node = *n
	}

	// add node to dst's children
	dst.Children = append(dst.Children, node)

	// return moved node
	return node, nil
}
