-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: CreateTransfers :one
INSERT INTO transfers (
  from_account_id, to_account_id, currency, amount
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteTransfers :exec
DELETE FROM transfers
WHERE id = $1;

-- name: UpdateTransfersAmount :one
UPDATE transfers 
SET amount = $2
WHERE id = $1
RETURNING *;