package app

import (
	"pjm.dev/sfs/graph"
)

func (a *app) GetNodeByURI(user graph.User, uri string) (graph.Node, error) {
	panic("not implemented")
}

func (a *app) RenameNode(user graph.User, id string, name string) (graph.Node, error) {
	panic("not implemented")
}

func (a *app) MoveNode(user graph.User, id string, dstID string) (graph.Node, error) {
	panic("not implemented")
}

func (a *app) GetRoot(user graph.User) (graph.Folder, error) {
	panic("not implemented")
}

func (a *app) InsertFolder(user graph.User, folder graph.Folder) (graph.Folder, error) {
	panic("not implemented")
}

func (a *app) GetFolderByID(user graph.User, id string) (graph.Folder, error) {
	panic("not implemented")
}

func (a *app) InsertFile(user graph.User, file graph.File) (graph.File, error) {
	panic("not implemented")
}

func (a *app) WriteFile(user graph.User, fileID string, content string) (graph.File, error) {
	panic("not implemented")
}

func (a *app) GetUserByID(id string) (graph.User, error) {
	panic("not implemented")
}
