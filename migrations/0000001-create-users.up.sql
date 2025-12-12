-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id            BIGSERIAL PRIMARY KEY,
    email         VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255),                    -- NULL for Google-only users
    google_id     VARCHAR(255) UNIQUE,
    name          VARCHAR(100) NOT NULL DEFAULT '',
    avatar_url    TEXT,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW()
);
