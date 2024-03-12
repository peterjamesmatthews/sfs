package mem

import (
	"errors"
	"fmt"

	"pjm.dev/sfs/graph"
	"pjm.dev/sfs/graph/model"
)

func (m *Database) InsertFile(user model.User, file model.File) (model.File, error) {
	parent, err := m.getFolderByID(file.Parent.ID)
	if errors.Is(err, graph.ErrNotFound) {
		return model.File{}, err
	} else if err != nil {
		return model.File{}, fmt.Errorf("failed to get parent %s: %w", file.Parent.ID, err)
	}

	// TODO verify user has write access to parent

	parent.Children = append(parent.Children, file)

	return file, nil
}

func (m *Database) GetFileByID(user model.User, id string) (model.File, error) {
	file, err := m.getFileByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return model.File{}, graph.ErrNotFound
	} else if err != nil {
		return model.File{}, fmt.Errorf("failed to get file %s: %w", id, err)
	}

	// TODO verify user has read access to file

	return *file, nil
}

func (m *Database) WriteFile(user model.User, fileID string, content string) (model.File, error) {
	// get file
	file, err := m.getFileByID(fileID)
	if errors.Is(err, graph.ErrNotFound) {
		return model.File{}, graph.ErrNotFound
	} else if err != nil {
		return model.File{}, fmt.Errorf("failed to get file %s: %w", fileID, err)
	}

	// TODO verify user has write access to file

	// write file's content
	file.Content = content

	// return written file
	return *file, nil
}

func (m *Database) getFileByID(id string) (*model.File, error) {
	node, err := m.getNodeByID(id)
	if errors.Is(err, graph.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to get file %s: %w", id, err)
	}

	file, ok := node.(*model.File)
	if !ok {
		return nil, fmt.Errorf("node %s is not a file", id)
	}

	return file, nil
}

func (m *Database) getFileByName(name string) (*model.File, error) {
	var file *model.File

	// get file by name
	node, err := m.getNodeByName(name)

	if errors.Is(err, graph.ErrNotFound) { // node wasn't found, create file
		file = &model.File{
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
	file, ok := node.(*model.File)
	if !ok {
		return nil, fmt.Errorf("node named %s is not a file", name)
	} else if file == nil {
		return nil, fmt.Errorf("nil file named %s", name)
	}

	return file, nil
}

func (m *Database) insertFile(file *model.File) (*model.File, error) {
	file.Parent.Children = append(file.Parent.Children, file)
	return file, nil
}
