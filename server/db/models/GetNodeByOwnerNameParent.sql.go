// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: GetNodeByOwnerNameParent.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getNodeByOwnerNameParent = `-- name: GetNodeByOwnerNameParent :one
SELECT id, owner, name, parent
FROM node
WHERE owner = $1
  AND name = $2
  AND (
    parent = $3
    OR (
      $3 IS NULL
      AND parent IS NULL
    )
  )
LIMIT 1
`

type GetNodeByOwnerNameParentParams struct {
	Owner  pgtype.UUID
	Name   string
	Parent pgtype.UUID
}

func (q *Queries) GetNodeByOwnerNameParent(ctx context.Context, arg GetNodeByOwnerNameParentParams) (Node, error) {
	row := q.db.QueryRow(ctx, getNodeByOwnerNameParent, arg.Owner, arg.Name, arg.Parent)
	var i Node
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Parent,
	)
	return i, err
}