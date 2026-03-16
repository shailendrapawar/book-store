CREATE TABLE IF NOT EXISTS users (
    id          VARCHAR(50)     PRIMARY KEY,
    name        VARCHAR(100)    NOT NULL,
    email       VARCHAR(150)    NOT NULL UNIQUE,
    password    VARCHAR(150)            NOT NULL,
    role        VARCHAR(20)     NOT NULL DEFAULT 'user',
    created_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);