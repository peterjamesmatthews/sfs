package app

import (
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var errForbidden = errors.New("forbidden")

func (a *App) isNotFoundError(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

func (a *App) isConflictError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgerrcode.IsIntegrityConstraintViolation(pgErr.Code)
	}

	return false
}
