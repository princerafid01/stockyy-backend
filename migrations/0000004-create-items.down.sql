-- +migrate Down
DROP INDEX IF EXISTS idx_items_user_search;
DROP INDEX IF EXISTS idx_items_user_barcode;
DROP INDEX IF EXISTS idx_items_user_location;
DROP INDEX IF EXISTS idx_items_user_category;
DROP INDEX IF EXISTS idx_items_user_quantity;
DROP TABLE IF EXISTS items;

