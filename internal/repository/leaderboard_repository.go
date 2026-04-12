package repository

import "context"

// LeaderboardEntry is a query result type, not a stored entity.
type LeaderboardEntry struct {
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	TotalScore int    `json:"total_score"`
	GamesWon   int    `json:"games_won"`
	Placement  *int   `json:"placement,omitempty"`
}

type LeaderboardRepository interface {
	GetHanchanLeaderboard(ctx context.Context, hanchanID int) ([]LeaderboardEntry, error)
	GetGroupLeaderboard(ctx context.Context, groupID int) ([]LeaderboardEntry, error)
}
