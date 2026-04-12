package repository

import (
	"context"
	"fmt"

	"HanchanManager/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HanchanRepository interface {
	Create(ctx context.Context, hanchan *domain.Hanchan) error
	GetByID(ctx context.Context, id int) (*domain.Hanchan, error)
	// ListByGroup(ctx context.Context, groupID int) ([]*domain.Hanchan, error)
	// AddPlayer(ctx context.Context, hp *domain.HanchanPlayer) error
	// ListPlayers(ctx context.Context, hanchanID int) ([]*domain.HanchanPlayer, error)
	// Close(ctx context.Context, hanchanID int, results []domain.HanchanPlayer) error
}

type hanchanRepo struct {
	db *pgxpool.Pool
}

func NewHanchanRepo(db *pgxpool.Pool) HanchanRepository {
	return &hanchanRepo{db: db}
}

func (r *hanchanRepo) Create(ctx context.Context, hanchan *domain.Hanchan) error {
	query := `INSERT INTO hanchans (group_id, name, date, uma) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(ctx, query, hanchan.GroupID, hanchan.Name, hanchan.Date, hanchan.Uma).Scan(&hanchan.ID)
}

func (r *hanchanRepo) GetByID(ctx context.Context, id int) (*domain.Hanchan, error) {
	h := &domain.Hanchan{}

	err := r.db.QueryRow(ctx,
		`SELECT id, group_id, name, date, status, uma, created_at FROM hanchans WHERE id = $1`,
		id,
	).Scan(&h.ID, &h.GroupID, &h.Name, &h.Date, &h.Status, &h.Uma, &h.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("GetHanchanByID: %w", err)
	}

	return h, nil
}
