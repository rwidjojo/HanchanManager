package repository

import (
	"context"
	"errors"
	"fmt"

	"HanchanManager/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepository interface {
	Create(ctx context.Context, group *domain.Group) error
	GetByID(ctx context.Context, id int) (*domain.Group, error)
	List(ctx context.Context) ([]*domain.Group, error)
}

type groupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) GroupRepository {
	return &groupRepo{db: db}
}

func (r *groupRepo) Create(ctx context.Context, group *domain.Group) error {
	query := `INSERT INTO groups (code, description) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRow(ctx, query, group.Code, group.Description).Scan(&group.ID, &group.CreatedAt)
}

func (r *groupRepo) GetByID(ctx context.Context, id int) (*domain.Group, error) {
	g := &domain.Group{}

	err := r.db.QueryRow(ctx,
		`SELECT id, code, description, created_at FROM groups WHERE id = $1`,
		id,
	).Scan(&g.ID, &g.Code, &g.Description, &g.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("get group: %w", err)
	}

	return g, nil
}

func (r *groupRepo) List(ctx context.Context) ([]*domain.Group, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, code, description, created_at FROM groups ORDER BY created_at`,
	)
	if err != nil {
		return nil, fmt.Errorf("list groups: %w", err)
	}
	defer rows.Close()

	var groups []*domain.Group
	for rows.Next() {
		g := &domain.Group{}
		if err := rows.Scan(&g.ID, &g.Code, &g.Description, &g.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan groups: %w", err)
		}
		groups = append(groups, g)
	}
	return groups, rows.Err()
}
