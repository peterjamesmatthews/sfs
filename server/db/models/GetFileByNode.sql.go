// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: GetFileByNode.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const getFileByNode = `-- name: GetFileByNode :one
SELECT id, node, content
FROM public.file
WHERE "node" = $1
LIMIT 1
`

func (q *Queries) GetFileByNode(ctx context.Context, node uuid.UUID) (File, error) {
	row := q.db.QueryRowContext(ctx, getFileByNode, node)
	var i File
	err := row.Scan(&i.ID, &i.Node, &i.Content)
	return i, err
}
