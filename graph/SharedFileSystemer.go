package graph

import "pjm.dev/sfs/graph/model"

type SharedFileSystemer interface {
	GetRoot() (model.Folder, error)
	GetNodeByID(id string) (model.Node, error)
	GetFolderByID(id string) (model.Folder, error)
	InsertFolder(folder model.Folder) (model.Folder, error)
}
