CREATE TABLE hanchans (
    id         INTEGER     PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    group_id   INTEGER     NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    name       TEXT,
    date       DATE        NOT NULL,
    status     TEXT        NOT NULL DEFAULT 'OPEN' CHECK (status IN ('OPEN', 'CLOSED')),
    uma        INT[]       NOT NULL DEFAULT '{15000,5000,-5000,-15000}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE hanchan_players (
    hanchan_id   INTEGER   NOT NULL REFERENCES hanchans(id) ON DELETE CASCADE,
    player_id    INTEGER   NOT NULL REFERENCES players(id)  ON DELETE RESTRICT,
    initial_seat TEXT      NOT NULL CHECK (initial_seat IN ('East', 'South', 'West', 'North')),
    final_score  INT,
    placement    INT       CHECK (placement BETWEEN 1 AND 4),
    PRIMARY KEY (hanchan_id, player_id)
);

CREATE INDEX idx_hanchans_group_id ON hanchans(group_id);
