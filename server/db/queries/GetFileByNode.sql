-- name: GetFileByNode :one
SELECT *
FROM public.file
WHERE "node" = $1
LIMIT 1;
