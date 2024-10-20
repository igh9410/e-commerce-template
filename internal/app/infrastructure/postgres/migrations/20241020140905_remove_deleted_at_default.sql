-- +goose Up
-- +goose StatementBegin

-- Alter the products table to make deleted_at nullable and remove the default value
ALTER TABLE products
  ALTER COLUMN deleted_at DROP DEFAULT,
  ALTER COLUMN deleted_at DROP NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Revert the change by setting deleted_at to NOT NULL and adding the default value
ALTER TABLE products
  ALTER COLUMN deleted_at SET DEFAULT NOW(),
  ALTER COLUMN deleted_at SET NOT NULL;

-- +goose StatementEnd
