package domain

import (
	"encoding/json"
	"time"
)

type RoundWind string

const (
	RoundWindEast  RoundWind = "EAST"
	RoundWindSouth RoundWind = "SOUTH"
	RoundWindWest  RoundWind = "WEST"
	RoundWindNorth RoundWind = "NORTH"
)

type GameOutcome string

const (
	GameOutcomeTsumo     GameOutcome = "TSUMO"
	GameOutcomeRon       GameOutcome = "RON"
	GameOutcomeRyuukyoku GameOutcome = "RYUUKYOKU"
	GameOutcomeChombo    GameOutcome = "CHOMBO"
)

type Game struct {
	ID                   int         `json:"id"`
	HanchanID            int         `json:"hanchan_id"`
	RoundWind            RoundWind   `json:"round_wind"`
	RoundNumber          int         `json:"round_number"`
	Honba                int         `json:"honba"`
	RiichiSticksCarried  int         `json:"riichi_sticks_carried"`
	RiichiSticksDeclared int         `json:"riichi_sticks_declared"`
	Outcome              GameOutcome `json:"outcome,omitempty"`
	CreatedAt            time.Time   `json:"created_at"`
}

// RiichiSticksTotal is derived — not stored in DB.
func (g *Game) RiichiSticksTotal() int {
	return g.RiichiSticksCarried + g.RiichiSticksDeclared
}

type PlayerRole string

const (
	PlayerRoleWinnerTsumo  PlayerRole = "WINNER_TSUMO"
	PlayerRoleWinnerRon    PlayerRole = "WINNER_RON"
	PlayerRoleDiscarder    PlayerRole = "DISCARDER"
	PlayerRoleNonDiscarder PlayerRole = "NON_DISCARDER"
	PlayerRoleTenpai       PlayerRole = "TENPAI"
	PlayerRoleNoten        PlayerRole = "NOTEN"
	PlayerRoleChombo       PlayerRole = "CHOMBO"
)

type GameResult struct {
	ID             int             `json:"id"`
	GameID         int             `json:"game_id"`
	PlayerID       int             `json:"player_id"`
	Role           PlayerRole      `json:"role"`
	RiichiDeclared bool            `json:"riichi_declared"`
	ScoreDelta     int             `json:"score_delta"`
	WinningHand    json.RawMessage `json:"winning_hand,omitempty"` // null for non-winners
}

// WinningHand is the structure stored as jsonb on the winner's GameResult row.
type WinningHand struct {
	Tiles           []string       `json:"tiles"`
	WinningTile     string         `json:"winning_tile"`
	WinType         string         `json:"win_type"` // "tsumo" | "ron"
	IsOpen          bool           `json:"is_open"`
	Yaku            []YakuEntry    `json:"yaku"`
	Han             int            `json:"han"`
	Fu              int            `json:"fu"`
	Dora            int            `json:"dora"`
	AkaDora         int            `json:"aka_dora"`
	Honba           int            `json:"honba,omitempty"`
	Payment         map[string]int `json:"payment"`
	SticksCarried   int            `json:"sticks_carried"`
	SticksDeclaried int            `json:"sticks_declared"`
	OwnRiichiStick  bool           `json:"own_riichi_stick"`
	NetDelta        int            `json:"net_delta"`
}

type YakuEntry struct {
	Name string `json:"name"`
	Han  int    `json:"han"`
}
