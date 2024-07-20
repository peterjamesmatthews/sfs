-- name: CreateNode :one
INSERT INTO node (owner, name, parent)
VALUES ($1, $2, $3)
RETURNING *;
