-- name: CreateUser :one
INSERT INTO "user" (name) VALUES ($1) RETURNING *;
