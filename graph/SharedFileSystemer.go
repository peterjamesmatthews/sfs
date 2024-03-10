package graph

import "pjm.dev/sfs/graph/model"

type SharedFileSystemer interface {
	GetNodeByID(id string) (model.Node, error)
	RenameNode(id string, name string) (model.Node, error)
	MoveNode(id string, dstID string) (model.Node, error)

	GetRoot() (model.Folder, error)
	InsertFolder(folder model.Folder) (model.Folder, error)
	GetFolderByID(id string) (model.Folder, error)

	InsertFile(file model.File) (model.File, error)
	GetFileByID(id string) (model.File, error)
	WriteFile(fileID string, content string) (model.File, error)
}
