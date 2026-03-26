package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rwidjojo/HanchanManager/internal/domain"
)

type PlayerRepository interface {
	Create(ctx context.Context, player *domain.Player) error
	// GetByID(ctx context.Context, id uuid.UUID) (*domain.Player, error)
	// List(ctx context.Context) ([]*domain.Player, error)
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

// func (r *playerRepo) GetByID(ctx context.Context, player *domain.Player) error {
// 	query := `INSERT INTO players (username, name) VALUES ($1, $2) RETURNING id`
// 	return r.db.QueryRow(ctx, query, player.Username, player).Scan(&player.ID)
// }

// func (r *playerRepo) Delete(ctx context.Context, id int64) error {
// 	query := `DELETE FROM players WHERE id = $1`
// 	_, err := r.db.Exec(ctx, query, id)
// 	return err
// }
