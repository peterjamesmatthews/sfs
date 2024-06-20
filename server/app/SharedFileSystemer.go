package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
		return graph.User{}, fmt.Errorf("failed to get user from token: %w", err)
	}

	// return user
	return a.getGraphUser(user), nil
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
	user, err := a.queries.CreateUser(
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
	return a.getGraphUser(user), nil
}

func (a *App) GetTokens(name string, password string) (graph.Tokens, error) {
	// get user by name
	user, err := a.queries.GetUserByName(context.Background(), name)
	if errors.Is(err, pgx.ErrNoRows) {
		return graph.Tokens{}, graph.ErrUnauthorized
	} else if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to get user: %w", err)
	}

	// compare password and user's salt with hash
	ok, err := a.comparePasswordWithSalt(password, user.Salt, user.Hash)
	if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to compare password: %w", err)
	} else if !ok {
		return graph.Tokens{}, graph.ErrUnauthorized
	}

	// generate access and refresh tokens for user
	access, refresh, err := a.generateTokensForUser(user)
	if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// convert and return tokens
	return a.getGraphTokens(access, refresh), nil
}

func (a *App) GetTokensFromAuth0Token(token string) (graph.Tokens, error) {
	// get user's name from Auth0 token
	id, name, err := a.getIDAndNameFromToken(token)
	if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to get user name from token: %w", err)
	}

	// get user by name
	var user models.User
	user, err = a.queries.GetUserByName(context.Background(), name)

	// create user if not found
	if errors.Is(err, pgx.ErrNoRows) {
		user, err = a.queries.CreateUser(
			context.Background(),
			models.CreateUserParams{
				Name:    name,
				Auth0ID: pgtype.Text{String: id, Valid: true},
			},
		)
	}

	if err != nil {
		return graph.Tokens{}, fmt.Errorf("failed to get user: %w", err)
	}

	// generate access and refresh tokens for user
	access, refresh, err := a.generateTokensForUser(user)
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
