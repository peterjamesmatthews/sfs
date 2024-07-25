// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: GetCountOfUsers.sql

package models

import (
	"context"
)

const getCountOfUsers = `-- name: GetCountOfUsers :one
SELECT COUNT(*)
FROM public.user
`

func (q *Queries) GetCountOfUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCountOfUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}
