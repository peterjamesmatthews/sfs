package graph

import "pjm.dev/sfs/graph/model"

type SharedFileSystemer interface {
	GetNodeByID(id string) (model.Node, error)
	RenameNode(model.Node) (model.Node, error)

	GetRoot() (model.Folder, error)
	InsertFolder(folder model.Folder) (model.Folder, error)
	GetFolderByID(id string) (model.Folder, error)

	InsertFile(file model.File) (model.File, error)
	GetFileByID(id string) (model.File, error)
	WriteFile(file model.File) (model.File, error)
}
