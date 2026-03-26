CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE players (
    id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    username   TEXT        UNIQUE,
    name       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE groups (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT        NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE group_members (
    group_id  UUID        NOT NULL REFERENCES groups(id)  ON DELETE CASCADE,
    player_id UUID        NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    role      TEXT        NOT NULL CHECK (role IN ('owner', 'admin', 'member')),
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (group_id, player_id)
);