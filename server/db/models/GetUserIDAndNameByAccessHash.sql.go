// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: GetUserIDAndNameByAccessHash.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getUserIDAndNameByAccessHash = `-- name: GetUserIDAndNameByAccessHash :one
SELECT public.user.id,
  public.user.name
FROM public.user
  JOIN access_token on owner = public.user.id
WHERE access_token.hash = $1
  AND expiration > CURRENT_TIMESTAMP
LIMIT 1
`

type GetUserIDAndNameByAccessHashRow struct {
	ID   pgtype.UUID
	Name string
}

func (q *Queries) GetUserIDAndNameByAccessHash(ctx context.Context, hash []byte) (GetUserIDAndNameByAccessHashRow, error) {
	row := q.db.QueryRow(ctx, getUserIDAndNameByAccessHash, hash)
	var i GetUserIDAndNameByAccessHashRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
