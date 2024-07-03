package app

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"pjm.dev/sfs/db/models"
)

func (a *App) isConflictError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == models.UniqueViolation
}

func (a *App) isNotFoundError(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
