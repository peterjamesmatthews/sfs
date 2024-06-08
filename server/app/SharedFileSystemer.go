package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"pjm.dev/sfs/db/models"
	"pjm.dev/sfs/graph"
)

func (a *App) Authenticate(auth string) (graph.User, error) {
	// get token from auth
	token := a.getTokenFromAuthorization(auth)

	// get hash of token
	hash := a.hashToken(token)

	// get user by hash
	userIDandName, err := a.q.GetUserIDAndNameByAccessHash(context.Background(), hash)
	if err != nil {
		return graph.User{}, fmt.Errorf("failed to get user by access hash: %w", err) // TODO handle pg.ErrNoRows
	}

	// return user
	return a.getGraphUser(models.User{ID: userIDandName.ID, Name: userIDandName.Name}), nil
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
	return a.getGraphUser(user), nil
}

func (a *App) GetTokens(name string, password string) (*graph.Tokens, error) {
	// get user by name
	user, err := a.q.GetUserByName(context.Background(), name)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == models.UniqueViolation {
		return nil, graph.ErrUnauthorized
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// compare password and user's salt with hash
	ok, err := a.comparePasswordWithSalt(password, user.Salt, user.Hash)
	if err != nil {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	} else if !ok {
		return nil, graph.ErrUnauthorized
	}

	// generate access and refresh tokens
	accessToken, refreshToken, err := a.generateTokensForUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// hash access token with sha256
	accessHash := a.hashToken(accessToken)

	// hash refresh token
	refreshHash := a.hashToken(refreshToken)

	// begin transaction for inserting tokens
	tx, err := a.db.Begin(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(context.Background())

	// insert access token
	_, err = a.q.WithTx(tx).InsertAccessToken(
		context.Background(),
		models.InsertAccessTokenParams{
			Owner: user.ID,
			Hash:  accessHash,
			// set expiration to 5 minutes from now
			Expiration: pgtype.Timestamp{Time: time.Now().Add(time.Minute * 5), Valid: true},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert access token: %w", err)
	}

	// insert refresh token
	_, err = a.q.WithTx(tx).InsertRefreshToken(
		context.Background(),
		models.InsertRefreshTokenParams{
			Owner: user.ID,
			Hash:  refreshHash,
			// set expiration to 1 week from now
			Expiration: pgtype.Timestamp{Time: time.Now().Add(time.Hour * 24 * 7), Valid: true},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert refresh token: %w", err)
	}

	// convert and return tokens
	tokens := a.getGraphTokens(accessToken, refreshToken)
	return &tokens, tx.Commit(context.Background())
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
