-- name: InsertAccessToken :one
INSERT INTO access_token (owner, hash, expiration)
VALUES ($1, $2, $3)
RETURNING *;
