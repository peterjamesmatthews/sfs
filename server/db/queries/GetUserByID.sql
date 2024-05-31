-- name: GetUserByID :one
SELECT *
FROM public.user
WHERE "id" = $1
LIMIT 1;
