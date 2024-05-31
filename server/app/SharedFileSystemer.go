package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"pjm.dev/sfs/db/models"
	"pjm.dev/sfs/graph"
)

func (a *App) Authenticate(auth string) (graph.User, error) {
	return graph.User{}, errors.ErrUnsupported
}

func (a *App) CreateUser(name string, password string) (graph.User, error) {
	// validate name
	if !a.isValidUserName(name) {
		return graph.User{}, fmt.Errorf("invalid user name: %s", name)
	}

	// generate salt
	salt, err := a.generateSalt()
	if err != nil {
		return graph.User{}, fmt.Errorf("failed to generate salt: %w", err)
	}

	// hash password with salt
	hash, err := a.hashPasswordWithSalt(password, salt)
	if err != nil {
		return graph.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// create user
	user, err := a.q.CreateUser(
		context.Background(),
		models.CreateUserParams{Name: name, Salt: salt, Hash: hash},
	)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == models.UniqueViolation {
		return graph.User{}, graph.ErrConflict
	} else if err != nil {
		return graph.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	// convert and return user
	return a.toGraphUser(user), nil
}

func (a *App) CreateFolder(creator graph.User, parentID *string, name string) (graph.Folder, error) {
	return graph.Folder{}, errors.ErrUnsupported
}

func (a *App) GetNodeByURI(user graph.User, uri string) (graph.Node, error) {
	return nil, errors.ErrUnsupported
}

func (a *App) RenameNode(user graph.User, id string, name string) (graph.Node, error) {
	return nil, errors.ErrUnsupported
}

func (a *App) MoveNode(user graph.User, id string, dstID *string) (graph.Node, error) {
	return nil, errors.ErrUnsupported
}

func (a *App) GetRoot(user graph.User) (graph.Folder, error) {
	return graph.Folder{}, errors.ErrUnsupported
}

func (a *App) InsertFolder(user graph.User, folder graph.Folder) (graph.Folder, error) {
	return graph.Folder{}, errors.ErrUnsupported
}

func (a *App) GetFolderByID(user graph.User, id string) (graph.Folder, error) {
	return graph.Folder{}, errors.ErrUnsupported
}

func (a *App) InsertFile(user graph.User, file graph.File) (graph.File, error) {
	return graph.File{}, errors.ErrUnsupported
}

func (a *App) WriteFile(user graph.User, fileID string, content string) (graph.File, error) {
	return graph.File{}, errors.ErrUnsupported
}

func (a *App) GetUserByID(id string) (graph.User, error) {
	return graph.User{}, errors.ErrUnsupported
}
