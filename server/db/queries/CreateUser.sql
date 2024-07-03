-- name: CreateUser :one
INSERT INTO public.user (email, auth0_id)
VALUES ($1, $2)
RETURNING *;
