package repository

import (
	"context"
	"fmt"

	"HanchanManager/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MembershipRepository interface {
	AddPlayer(ctx context.Context, groupID int, playerID int) error
	GetPlayers(ctx context.Context, groupID int) ([]*domain.Player, error)
}

type membershipRepo struct {
	db *pgxpool.Pool
}

func NewMembershipRepo(db *pgxpool.Pool) MembershipRepository {
	return &membershipRepo{db: db}
}

func (r *membershipRepo) AddPlayer(ctx context.Context, groupID int, playerID int) error {
	query := `INSERT INTO group_members (group_id, player_id) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, groupID, playerID)
	return err
}

func (r *membershipRepo) GetPlayers(ctx context.Context, groupID int) ([]*domain.Player, error) {
	query := `
        SELECT p.id, p.username, p.name, p.created_at
        FROM players p
        JOIN group_members gm ON gm.player_id = p.id
        WHERE gm.group_id = $1
    `

	rows, err := r.db.Query(ctx, query, groupID)
	if err != nil {
		return nil, fmt.Errorf("get player: %w", err)
	}
	defer rows.Close()

	var players []*domain.Player
	for rows.Next() {
		p := &domain.Player{}
		if err := rows.Scan(&p.ID, &p.Username, &p.Name, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan players: %w", err)
		}
		players = append(players, p)
	}
	return players, nil
}
