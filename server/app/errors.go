package app

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var errForbidden = errors.New("forbidden")

func (a *App) isNotFoundError(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

func (a *App) isConflictError(err error) bool {
	return false // TODO implement
}
