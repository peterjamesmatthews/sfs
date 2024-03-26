package mem

import (
	"errors"
	"fmt"
	"slices"

	"pjm.dev/sfs/graph"
)

func (m *Database) GetNodeByID(user graph.User, id string) (graph.Node, error) {
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get node %s: %w", id, err)
	}

	// verify user has read access to node
	if hasAccess, err := m.has(user, graph.AccessTypeRead, node); err != nil {
		return nil, fmt.Errorf("failed to check user %s has %s access on node %s: %w", user.ID, graph.AccessTypeRead, node.GetID(), err)
	} else if !hasAccess {
		return nil, graph.ErrUnauthorized
	}

	return node, nil
}

func (m *Database) getNodeByID(id string) (graph.Node, error) {
	if id == m.Root.ID {
		return m.Root, nil
	}

	nodes := m.Root.Children
	for _, node := range nodes {
		if node.GetID() == id {
			return node, nil
		}

		folder, ok := node.(*graph.Folder)
		if !ok {
			continue
		}

		nodes = append(nodes, folder.Children...)
	}

	return nil, graph.ErrNotFound
}

func (m *Database) RenameNode(user graph.User, id string, name string) (graph.Node, error) {
	// get node
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get node %s: %w", id, err)
	}

	// verify user has write access to node
	if hasAccess, err := m.has(user, graph.AccessTypeWrite, node); err != nil {
		return nil, fmt.Errorf("failed to check user %s has %s access on node %s: %w", user.ID, graph.AccessTypeWrite, node.GetID(), err)
	} else if !hasAccess {
		return nil, graph.ErrUnauthorized
	}

	// set update node's name
	switch n := node.(type) {
	case *graph.Folder:
		n.Name = name
		node = *n
	case *graph.File:
		n.Name = name
		node = *n
	}

	// return renamed node
	return node, nil
}

func (m *Database) MoveNode(user graph.User, id string, dstID string) (graph.Node, error) {
	// get node
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to rename node %s: %w", node.GetID(), err)
	}

	// verify user owns the node
	if owns := m.owns(user, node); !owns {
		return nil, graph.ErrUnauthorized
	}

	// get destination parent folder
	dst, err := m.getFolderByID(dstID)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to rename node %s: %w", node.GetID(), err)
	}

	// verify user has write access to destination parent
	if hasAccess, err := m.has(user, graph.AccessTypeWrite, node); err != nil {
		return nil, fmt.Errorf("failed to check user %s has %s access on parent %s: %w", user.ID, graph.AccessTypeWrite, dst.ID, err)
	} else if !hasAccess {
		return nil, graph.ErrUnauthorized
	}

	// get source parent folder
	src := node.GetParent()

	// remove node from src's children
	src.Children = slices.DeleteFunc(src.Children, func(child graph.Node) bool {
		return child.GetID() == node.GetID()
	})

	// set node's parent to dst
	switch n := node.(type) {
	case *graph.Folder:
		n.Parent = dst
		node = *n
	case *graph.File:
		n.Parent = dst
		node = *n
	}

	// add node to dst's children
	dst.Children = append(dst.Children, node)

	// return moved node
	return node, nil
}

// GetNodeByURI fetches a node by its uri.
//
// # Arguments
//   - user: The user who is fetching the node.
//   - uri: a '/' separated path of node names.
//
// # Errors
//   - `ErrNotFound` if the node is not found.
//   - `ErrUnauthorized` if `user` does not have read access to any of the nodes in the uri.
func (m *Database) GetNodeByURI(user graph.User, uri string) (graph.Node, error) {
	target := URI{uri: uri}
	names := target.GetNames()

	// check each name in the uri for read access, return the last node
	var current = RootURI
	var node graph.Node
	var err error
	for _, name := range names {
		if name == RootName {
			node = m.Root
		} else {
			current = current.AddNames(name)
			node, err = m.getNodeByURI(current)
			if err != nil {
				return nil, fmt.Errorf("failed to get node at uri %s: %w", uri, err)
			}
		}

		hasAccess, err := m.has(user, graph.AccessTypeRead, node)
		if err != nil {
			return nil, fmt.Errorf("failed to check if user %s has %s access to node at uri %s: %w", user.ID, graph.AccessTypeRead, uri, err)
		} else if !hasAccess {
			return nil, graph.ErrUnauthorized
		}
	}

	if node == nil {
		return nil, graph.ErrNotFound
	}

	return node, nil
}

func (m *Database) getNodeByURI(uri URI) (graph.Node, error) {
	if !uri.IsAbsolute() {
		return nil, errors.ErrUnsupported
	}

	names := uri.GetNames()
	names = names[1:] // remove root name

	// start from root
	var node graph.Node
	folder := m.Root
	for i, name := range names {
		for _, child := range folder.Children {
			if child.GetName() != name {
				continue
			}

			if i == len(names)-1 {
				return child, nil
			}

			if _, ok := child.(*graph.File); ok {
				continue
			}

			folder = child.(*graph.Folder)
		}
	}

	if node == nil {
		return nil, graph.ErrNotFound
	}

	return node, nil
}
