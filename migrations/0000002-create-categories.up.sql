-- +migrate Up
CREATE TABLE IF NOT EXISTS categories (
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name       VARCHAR(80)  NOT NULL,
    color      VARCHAR(7)   NOT NULL DEFAULT '#10B981',  -- emerald green default
    created_at TIMESTAMPTZ  DEFAULT NOW(),
    updated_at TIMESTAMPTZ  DEFAULT NOW(),
    UNIQUE(user_id, name)
);

