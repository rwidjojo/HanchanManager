CREATE TABLE players (
    id         INTEGER     PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username   TEXT        UNIQUE NOT NULL,
    name       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE groups (
    id          INTEGER     PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    code        TEXT        UNIQUE NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE group_members (
    group_id  INTEGER     NOT NULL REFERENCES groups(id)  ON DELETE CASCADE,
    player_id INTEGER     NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    role      TEXT        NOT NULL CHECK (role IN ('OWNER', 'ADMIN', 'MEMBER')),
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (group_id, player_id)
);
