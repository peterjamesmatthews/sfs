package app

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"pjm.dev/sfs/graph"
)

func (a *App) Authenticate(r *http.Request) (graph.User, error) {
	// get Authorization header
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return graph.User{}, graph.ErrUnauthorized
	}

	// parse auth to user id
	id := pgtype.UUID{}
	err := id.Scan(auth)
	if err != nil {
		return graph.User{}, graph.ErrUnauthorized
	}

	// get user by id
	user, err := a.q.GetUserByID(r.Context(), id)
	if err != nil {
		return graph.User{}, graph.ErrUnauthorized
	}

	// convert user to graph user
	return a.toGraphUser(user), nil
}

func (a *App) CreateUser(name string) (graph.User, error) {
	return graph.User{}, errors.New("not implemented")
}

func (a *App) CreateFolder(creator graph.User, parentID *string, name string) (graph.Folder, error) {
	return graph.Folder{}, errors.New("not implemented")
}

func (a *App) GetNodeByURI(user graph.User, uri string) (graph.Node, error) {
	return nil, errors.New("not implemented")
}

func (a *App) RenameNode(user graph.User, id string, name string) (graph.Node, error) {
	return nil, errors.New("not implemented")
}

func (a *App) MoveNode(user graph.User, id string, dstID *string) (graph.Node, error) {
	return nil, errors.New("not implemented")
}

func (a *App) GetRoot(user graph.User) (graph.Folder, error) {
	return graph.Folder{}, errors.New("not implemented")
}

func (a *App) InsertFolder(user graph.User, folder graph.Folder) (graph.Folder, error) {
	return graph.Folder{}, errors.New("not implemented")
}

func (a *App) GetFolderByID(user graph.User, id string) (graph.Folder, error) {
	return graph.Folder{}, errors.New("not implemented")
}

func (a *App) InsertFile(user graph.User, file graph.File) (graph.File, error) {
	return graph.File{}, errors.New("not implemented")
}

func (a *App) WriteFile(user graph.User, fileID string, content string) (graph.File, error) {
	return graph.File{}, errors.New("not implemented")
}

func (a *App) GetUserByID(id string) (graph.User, error) {
	return graph.User{}, errors.New("not implemented")
}
