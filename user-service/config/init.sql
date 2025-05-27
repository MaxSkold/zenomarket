-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    email      TEXT NOT NULL UNIQUE,
    full_name  TEXT NOT NULL,
    phone      TEXT,
    avatar_url TEXT,
    is_active  BOOLEAN     DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

ALTER TABLE users
    OWNER TO dev;

-- Создание таблицы учётных записей
CREATE TABLE IF NOT EXISTS auth_accounts
(
    id             SERIAL PRIMARY KEY,
    user_id        INTEGER NOT NULL UNIQUE REFERENCES users (id) ON DELETE CASCADE,
    password_hash  TEXT    NOT NULL,
    password_salt  TEXT    NOT NULL,
    provider       TEXT        DEFAULT 'local',
    third_party_id TEXT,
    last_login_at  TIMESTAMPTZ,
    updated_at     TIMESTAMPTZ DEFAULT now()
);

ALTER TABLE auth_accounts
    OWNER TO dev;
