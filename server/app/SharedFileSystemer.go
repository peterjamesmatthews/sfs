package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"pjm.dev/sfs/db/models"
	"pjm.dev/sfs/graph"
)

func (a *App) Authenticate(auth string) (graph.User, error) {
	// get token from auth
	token := a.getTokenFromAuthorization(auth)

	// get user from token
	user, err := a.getUserFromToken(token)
	if err != nil {
		return graph.User{}, graph.ErrUnauthorized
	}

	// return user
	return a.getGraphUser(user), nil
}

func (a *App) GetTokensFromAuth0Token(token string) (graph.Tokens, error) {
	// get user's email from Auth0 token
	id, email, err := a.Auth0.GetIDAndEmailFromToken(token)
	if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to get user name from token: %w", err)
	}

	// get user by email
	var user models.User
	user, err = a.Queries.GetUserByEmail(context.Background(), email)

	// create user if not found
	if a.isNotFoundError(err) {
		user, err = a.Queries.CreateUser(
			context.Background(),
			models.CreateUserParams{
				Email:   email,
				Auth0ID: pgtype.Text{String: id, Valid: true},
			},
		)
	}

	if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to get user: %w", err)
	}

	// generate access and refresh tokens for user
	access, refresh, err := a.GetTokensForUser(user)
	if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// convert and return tokens
	return a.getGraphTokens(access, refresh), nil
}

func (a *App) RefreshTokens(refresh string) (graph.Tokens, error) {
	access, refresh, err := a.refreshTokens(refresh)
	if errors.Is(err, errExpired) {
		return graph.Tokens{}, graph.ErrForbidden
	} else if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to refresh tokens: %w", err)
	}
	return a.getGraphTokens(access, refresh), nil
}

func (a *App) CreateFolder(gUser graph.User, path string) (graph.Folder, error) {
	// get user
	user, err := a.getUserFromGraphUser(gUser)
	if a.isNotFoundError(err) {
		return graph.Folder{}, graph.ErrUnauthorized
	} else if err != nil {
		return graph.Folder{}, fmt.Errorf("failed to get user: %w", err)
	}

	// create node
	node, err := a.createNode(user, path)
	if a.isNotFoundError(err) {
		return graph.Folder{}, graph.ErrNotFound
	} else if a.isConflictError(err) {
		return graph.Folder{}, graph.ErrConflict
	} else if errors.Is(err, errForbidden) {
		return graph.Folder{}, graph.ErrForbidden
	} else if err != nil {
		return graph.Folder{}, fmt.Errorf("failed to create folder: %w", err)
	}

	// return created node as folder
	return a.getGraphFolder(node), nil
}

func (a *App) GetNodeFromPath(gUser graph.User, path string) (graph.Node, error) {
	// get user
	user, err := a.getUserFromGraphUser(gUser)
	if a.isNotFoundError(err) {
		return nil, graph.ErrUnauthorized
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// get node from path
	node, err := a.getNodeFromPath(user, path)
	if a.isNotFoundError(err) {
		return nil, graph.ErrNotFound
	} else if errors.Is(err, errForbidden) {
		return nil, graph.ErrForbidden
	}

	// return converted node
	return a.getGraphNode(node)
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
