-- +migrate Up
-- Fix existing rows where image_urls contains '{}' (JSON object) instead of '[]' (JSON array)
-- This updates any TEXT[] column that might have been stored with invalid JSON object format

UPDATE items 
SET image_urls = '{}' 
WHERE image_urls::text = '{}';

-- Note: PostgreSQL empty array is represented as '{}' in PostgreSQL array syntax
-- The application will now handle this correctly
