-- name: GetUserByEmail :one
SELECT *
FROM public.user
WHERE "email" = $1
LIMIT 1;
