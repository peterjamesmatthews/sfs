package mem

import (
	"errors"
	"fmt"

	"pjm.dev/sfs/graph"
)

func (m *Database) InsertFile(user graph.User, file graph.File) (graph.File, error) {
	// get parent
	parent, err := m.getFolderByID(file.Parent.ID)
	if errors.Is(err, graph.ErrNotFound) {
		return graph.File{}, err
	} else if err != nil {
		return graph.File{}, fmt.Errorf("failed to get parent %s: %w", file.Parent.ID, err)
	}

	// verify user has write access on parent
	if hasAccess, err := m.has(user, graph.AccessTypeWrite, parent); err != nil {
		return graph.File{}, fmt.Errorf("failed to check user %s has %s access on parent %s: %w", user.ID, graph.AccessTypeWrite, parent.ID, err)
	} else if !hasAccess {
		return graph.File{}, graph.ErrUnauthorized
	}

	parent.Children = append(parent.Children, file)

	return file, nil
}

func (m *Database) GetFileByID(user graph.User, id string) (graph.File, error) {
	// get file
	file, err := m.getFileByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return graph.File{}, graph.ErrNotFound
	} else if err != nil {
		return graph.File{}, fmt.Errorf("failed to get file %s: %w", id, err)
	}

	// verify user has read access to file
	if hasAccess, err := m.has(user, graph.AccessTypeRead, file); err != nil {
		return graph.File{}, fmt.Errorf("failed to check user %s has %s access on file %s: %w", user.ID, graph.AccessTypeRead, file.ID, err)
	} else if !hasAccess {
		return graph.File{}, graph.ErrUnauthorized
	}

	return *file, nil
}

func (m *Database) WriteFile(user graph.User, fileID string, content string) (graph.File, error) {
	// get file
	file, err := m.getFileByID(fileID)
	if errors.Is(err, graph.ErrNotFound) {
		return graph.File{}, graph.ErrNotFound
	} else if err != nil {
		return graph.File{}, fmt.Errorf("failed to get file %s: %w", fileID, err)
	}

	// verify user has write access to file
	if hasAccess, err := m.has(user, graph.AccessTypeWrite, file); err != nil {
		return graph.File{}, fmt.Errorf("failed to check user %s has %s access on file %s: %w", user.ID, graph.AccessTypeWrite, file.ID, err)
	} else if !hasAccess {
		return graph.File{}, graph.ErrUnauthorized
	}

	// write file's content
	file.Content = content

	// return written file
	return *file, nil
}

func (m *Database) getFileByID(id string) (*graph.File, error) {
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get file %s: %w", id, err)
	}

	file, ok := node.(*graph.File)
	if !ok {
		return nil, fmt.Errorf("node %s is not a file", id)
	}

	return file, nil
}

func (m *Database) getFileByName(name string) (*graph.File, error) {
	var file *graph.File

	// get file by name
	node, err := m.getNodeByName(name)

	if errors.Is(err, graph.ErrNotFound) { // node wasn't found, create file
		file = &graph.File{
			ID:      m.Generate().String(),
			Name:    name,
			Owner:   nil,
			Parent:  m.Root,
			Content: "",
		}
		file, err = m.insertFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to insert file %s: %w", name, err)
		}
		return file, nil
	} else if err != nil { // error occurred getting node, bail
		return nil, fmt.Errorf("failed to get node named %s: %w", name, err)
	}

	// try casting node to file
	file, ok := node.(*graph.File)
	if !ok {
		return nil, fmt.Errorf("node named %s is not a file", name)
	} else if file == nil {
		return nil, fmt.Errorf("nil file named %s", name)
	}

	return file, nil
}

func (m *Database) insertFile(file *graph.File) (*graph.File, error) {
	file.Parent.Children = append(file.Parent.Children, file)
	return file, nil
}
