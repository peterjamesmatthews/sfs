-- name: InsertRefreshToken :one
INSERT INTO refresh_token (owner, hash, expiration)
VALUES ($1, $2, $3)
RETURNING *;
