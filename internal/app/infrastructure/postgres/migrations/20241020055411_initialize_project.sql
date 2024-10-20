-- +goose Up
-- +goose StatementBegin

-- Create products table
CREATE TABLE products (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(100),
    price BIGINT NOT NULL,
    status VARCHAR(100) DEFAULT 'active',  -- Use VARCHAR instead of ENUM
    tags TEXT[],
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NOW()
);

-- Create product_stocks table
CREATE TABLE product_stocks (
    id SERIAL PRIMARY KEY,
    product_id UUID NOT NULL,
    location TEXT,
    sold_at TIMESTAMPTZ,
    expiry_date DATE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS product_stocks;
DROP TABLE IF EXISTS products;

-- +goose StatementEnd
