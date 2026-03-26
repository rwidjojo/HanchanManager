package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rwidjojo/HanchanManager/internal/domain"
)

type PlayerRepository interface {
	Create(ctx context.Context, player *domain.Player) error
	GetByUsername(ctx context.Context, username string) (*domain.Player, error)
	List(ctx context.Context) ([]*domain.Player, error)
}

type playerRepo struct {
	db *pgxpool.Pool
}

func NewPlayerRepo(db *pgxpool.Pool) PlayerRepository {
	return &playerRepo{db: db}
}

func (r *playerRepo) Create(ctx context.Context, player *domain.Player) error {
	query := `INSERT INTO players (username, name) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(ctx, query, player.Username, player.Name).Scan(&player.ID)
}

func (r *playerRepo) GetByUsername(ctx context.Context, username string) (*domain.Player, error) {
	p := &domain.Player{}

	err := r.db.QueryRow(ctx,
		`SELECT id, username, name, created_at FROM players WHERE username = $1`,
		username,
	).Scan(&p.ID, &p.Username, &p.Name, &p.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("GetPlayerByID: %w", err)
	}

	return p, nil
}

func (r *playerRepo) List(ctx context.Context) ([]*domain.Player, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, username, name, created_at FROM players ORDER BY created_at`,
	)
	if err != nil {
		return nil, fmt.Errorf("ListPlayers: %w", err)
	}
	defer rows.Close()

	var players []*domain.Player
	for rows.Next() {
		p := &domain.Player{}
		if err := rows.Scan(&p.ID, &p.Username, &p.Name, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("ListPlayers scan: %w", err)
		}
		players = append(players, p)
	}
	return players, rows.Err()
}
