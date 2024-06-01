-- name: GetUserByName :one
SELECT *
FROM public.user
WHERE "name" = $1
LIMIT 1;
