CREATE TABLE games (
    id                      INTEGER     PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    hanchan_id              INTEGER     NOT NULL REFERENCES hanchans(id) ON DELETE CASCADE,
    round_wind              TEXT        NOT NULL CHECK (round_wind IN ('EAST', 'SOUTH', 'WEST', 'NORTH')),
    round_number            INT         NOT NULL CHECK (round_number BETWEEN 1 AND 4),
    honba                   INT         NOT NULL DEFAULT 0 CHECK (honba >= 0),
    riichi_sticks_carried   INT         NOT NULL DEFAULT 0 CHECK (riichi_sticks_carried >= 0),
    riichi_sticks_declared  INT         NOT NULL DEFAULT 0 CHECK (riichi_sticks_declared >= 0),
    outcome                 TEXT        NOT NULL CHECK (outcome IN ('TSUMO', 'RON', 'RYUUKYOKU', 'CHOMBO')),
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE game_results (
    id              INTEGER     PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    game_id         INTEGER     NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    player_id       INTEGER     NOT NULL REFERENCES players(id) ON DELETE RESTRICT,
    role            TEXT        NOT NULL CHECK (role IN (
                                'WINNER_TSUMO', 'WINNER_RON',
                                'DISCARDER', 'NON_DISCARDER',
                                'TENPAI', 'NOTEN', 'CHOMBO'
                            )),
    riichi_declared BOOLEAN NOT NULL DEFAULT FALSE,
    score_delta     INT     NOT NULL,
    winning_hand    JSONB,
    UNIQUE (game_id, player_id)
);

CREATE INDEX idx_games_hanchan_id       ON games(hanchan_id);
CREATE INDEX idx_game_results_game_id   ON game_results(game_id);
CREATE INDEX idx_game_results_player_id ON game_results(player_id);
