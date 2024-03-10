package memdb

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"pjm.dev/sfs/graph/model"
)

func (m *MemDatabase) InsertFile(file model.File) (model.File, error) {
	parent, err := m.getFolderByID(file.Parent.ID)
	if err != nil {
		return model.File{}, fmt.Errorf("failed to get parent %s", file.Parent.ID)
	}

	parent.Children = append(parent.Children, file)

	return file, nil
}

func (m *MemDatabase) GetFileByID(id string) (model.File, error) {
	file, err := m.getFileByID(id)
	if err != nil {
		return model.File{}, fmt.Errorf("failed to get file %s: %w", id, err)
	} else if file == nil {
		return model.File{}, fmt.Errorf("nil file %s", id)
	}
	return *file, nil
}

func (m *MemDatabase) WriteFile(fileID string, content string) (model.File, error) {
	// get file
	file, err := m.getFileByID(fileID)
	if err != nil {
		return model.File{}, fmt.Errorf("failed to get file %s", fileID)
	}

	// write file's content
	file.Content = content

	// return written file
	return *file, nil
}

func (m *MemDatabase) getFileByID(id string) (*model.File, error) {
	node, err := m.getNodeByID(id)
	if errors.Is(err, errNodeNotFound) {
		return nil, fmt.Errorf("file %s not found: %w", id, err)
	}

	file, ok := node.(*model.File)
	if !ok {
		return nil, fmt.Errorf("node %s is not a file", id)
	}

	return file, nil
}

func (m *MemDatabase) getFileByName(name string) (*model.File, error) {
	var file *model.File

	// get file by name
	node, err := m.getNodeByName(name)

	if errors.Is(err, errNodeNotFound) { // node wasn't found, create file
		file = &model.File{
			ID:      uuid.NewString(),
			Name:    name,
			Owner:   nil,
			Parent:  m.root,
			Content: "",
		}
		file, err = m.insertFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to create file %s: %w", name, err)
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

func (m *MemDatabase) insertFile(file *model.File) (*model.File, error) {
	file.Parent.Children = append(file.Parent.Children, file)
	return file, nil
}
