-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: CreateEntries :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEntries :exec
DELETE FROM entries
WHERE id = $1;

-- name: UpdateEntriesAmount :one
UPDATE entries 
SET amount = $2
WHERE id = $1
RETURNING *;