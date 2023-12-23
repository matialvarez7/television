-- name: CreateDevice :one
INSERT INTO devices (id, created_at, updated_at, codetag)
VALUES ($1, $2, $3, $4)
RETURNING *;