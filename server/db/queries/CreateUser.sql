-- name: CreateUser :one
INSERT INTO public.user (name, salt, hash, auth0_id)
VALUES ($1, $2, $3, $4)
RETURNING *;
