-- name: CreateUser :one
INSERT INTO public.user (name, salt, hash)
VALUES ($1, $2, $3)
RETURNING *;
