package app

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

func (a *App) isNotFoundError(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
