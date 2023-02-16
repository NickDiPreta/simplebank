-- name: CreateAccount :one
INSERT INTO accounts (OWNER, balance, currency)
  VALUES ($1, $2, $3)
RETURNING
  *;

