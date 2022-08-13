-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountsForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR No KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateAccounts :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteAccounts :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccountsBalance :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountsBalance :one
UPDATE accounts 
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;