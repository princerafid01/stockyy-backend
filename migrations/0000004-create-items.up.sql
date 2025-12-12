-- +migrate Up
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS items (
    id                  BIGSERIAL PRIMARY KEY,
    user_id             BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    name                VARCHAR(255) NOT NULL,
    description         TEXT,
    
    quantity            INTEGER      NOT NULL DEFAULT 0,
    low_stock_threshold INTEGER      NOT NULL DEFAULT 10,
    
    price               DECIMAL(12,2),
    cost_price          DECIMAL(12,2),                -- optional: for profit calc later
    
    sku                 VARCHAR(100),
    barcode             VARCHAR(100),
    
    category_id         BIGINT REFERENCES categories(id) ON DELETE SET NULL,
    location_id         BIGINT REFERENCES locations(id) ON DELETE SET NULL,
    
    image_urls          TEXT[]       DEFAULT '{}',    -- array of uploaded image URLs
    notes               TEXT,
    
    created_at          TIMESTAMPTZ  DEFAULT NOW(),
    updated_at          TIMESTAMPTZ  DEFAULT NOW()
);

-- Indexes for speed
CREATE INDEX idx_items_user_quantity     ON items(user_id, quantity);
CREATE INDEX idx_items_user_category     ON items(user_id, category_id);
CREATE INDEX idx_items_user_location     ON items(user_id, location_id);
CREATE INDEX idx_items_user_barcode      ON items(user_id, barcode);
CREATE INDEX idx_items_user_search       ON items(user_id, name gin_trgm_ops) USING GIN;  -- for fast fuzzy search

