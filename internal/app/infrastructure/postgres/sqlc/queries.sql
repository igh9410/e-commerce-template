-- name: CreateProduct :one
INSERT INTO products (name, description, category, price, status, tags, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
RETURNING id, name, description, category, price, status, tags, created_at, updated_at, deleted_at;
