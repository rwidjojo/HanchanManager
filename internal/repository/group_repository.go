package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rwidjojo/HanchanManager/internal/domain"
)

type GroupRepository interface {
	Create(ctx context.Context, group *domain.Group) error
	GetByID(ctx context.Context, id int) (*domain.Group, error)
}

type groupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) GroupRepository {
	return &groupRepo{db: db}
}

func (r *groupRepo) Create(ctx context.Context, group *domain.Group) error {
	query := `INSERT INTO groups (code, description) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(ctx, query, group.Code, group.Description).Scan(&group.ID)
}

func (r *groupRepo) GetByID(ctx context.Context, id int) (*domain.Group, error) {
	g := &domain.Group{}

	err := r.db.QueryRow(ctx,
		`SELECT id, code, description, created_at FROM groups WHERE id = $1`,
		id,
	).Scan(&g.ID, &g.Code, &g.Description, &g.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("GetGroupByID: %w", err)
	}

	return g, nil
}
