-- +migrate Down
DROP INDEX IF EXISTS idx_movements_user;
DROP INDEX IF EXISTS idx_movements_item;
DROP TABLE IF EXISTS stock_movements;

