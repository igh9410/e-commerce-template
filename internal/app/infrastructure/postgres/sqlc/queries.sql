-- name: CreateProduct :one
INSERT INTO products (
    id, name, description, category, price, status, tags, created_at, updated_at, deleted_at
) VALUES (
    gen_random_uuid(), $1, $2, $3, $4, $5, $6, NOW(), NOW(), NOW()
)
RETURNING id, name, description, category, price, status, tags, created_at, updated_at, deleted_at;
