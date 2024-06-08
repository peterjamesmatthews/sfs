-- name: GetUserIDAndNameByAccessHash :one
SELECT public.user.id,
  public.user.name
FROM public.user
  JOIN access_token on owner = public.user.id
WHERE access_token.hash = $1
  AND expiration > CURRENT_TIMESTAMP
LIMIT 1;
