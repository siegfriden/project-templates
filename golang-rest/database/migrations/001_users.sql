-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE FUNCTION uuid_zero() RETURNS UUID AS $$
  SELECT '00000000-0000-0000-0000-000000000000'::UUID;
$$ LANGUAGE sql IMMUTABLE;

CREATE TYPE user_status AS ENUM (
    'ACTIVE',
    'DEACTIVATED',
    'SUSPENDED'
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() CHECK (id != uuid_zero()),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    name TEXT NOT NULL CHECK (name != ''),
    status user_status NOT NULL DEFAULT 'ACTIVE',

    email CITEXT UNIQUE NOT NULL CHECK (email != ''),
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    email_verification_token TEXT,
    email_verification_expires TIMESTAMPTZ,

    password_hash TEXT NOT NULL,
    password_reset_token TEXT,
    password_reset_expires TIMESTAMPTZ,

    last_login TIMESTAMPTZ,
    last_password_change TIMESTAMPTZ
);

CREATE INDEX idx_users_name ON users USING gin(name gin_trgm_ops);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_email_verification_token ON users(email_verification_token);
CREATE INDEX idx_users_password_reset_token ON users(password_reset_token);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_password_reset_token;
DROP INDEX IF EXISTS idx_users_email_verification_token;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_name;

DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS user_status;

DROP FUNCTION IF EXISTS uuid_zero();

DROP EXTENSION IF EXISTS pg_trgm;
DROP EXTENSION IF EXISTS citext;
-- +goose StatementEnd
