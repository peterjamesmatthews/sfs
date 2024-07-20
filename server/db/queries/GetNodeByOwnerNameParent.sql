-- name: GetNodeByOwnerNameParent :one
SELECT *
FROM node
WHERE owner = $1
  AND name = $2
  AND (
    parent = $3
    OR (
      $3 IS NULL
      AND parent IS NULL
    )
  )
LIMIT 1;
