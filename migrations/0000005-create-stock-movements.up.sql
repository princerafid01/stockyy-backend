-- +migrate Up
CREATE TABLE IF NOT EXISTS stock_movements (
    id              BIGSERIAL PRIMARY KEY,
    item_id         BIGINT NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    user_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    quantity_change INTEGER NOT NULL,          -- positive = in, negative = out
    new_quantity    INTEGER NOT NULL,          -- quantity after this movement
    reason          VARCHAR(255),              -- "Sold", "Received", "Damage", "Manual adjust"…
    
    created_at      TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_movements_item   ON stock_movements(item_id);
CREATE INDEX idx_movements_user   ON stock_movements(user_id);

