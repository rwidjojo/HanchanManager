package repository

import (
	"context"
	"errors"

	"HanchanManager/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GameRepository interface {
	Create(ctx context.Context, game *domain.Game) error
	GetByID(ctx context.Context, id int) (*domain.Game, error)
	ListByHanchan(ctx context.Context, hanchanID int) ([]*domain.Game, error)
	SaveResults(ctx context.Context, results []*domain.GameResult) error
	GetResults(ctx context.Context, gameID int) ([]*domain.GameResult, error)
}

type gameRepo struct {
	db *pgxpool.Pool
}

func NewGameRepo(db *pgxpool.Pool) GameRepository {
	return &gameRepo{db: db}
}

func (r *gameRepo) Create(ctx context.Context, game *domain.Game) error {
	return errors.New("Method not yet implemented")
}

func (r *gameRepo) GetByID(ctx context.Context, id int) (*domain.Game, error) {
	return nil, errors.New("Method not yet implemented")
}

func (r *gameRepo) ListByHanchan(ctx context.Context, hanchanID int) ([]*domain.Game, error) {
	return nil, errors.New("Method not yet implemented")
}

func (r *gameRepo) SaveResults(ctx context.Context, results []*domain.GameResult) error {
	return errors.New("Method not yet implemented")
}

func (r *gameRepo) GetResults(ctx context.Context, gameID int) ([]*domain.GameResult, error) {
	return nil, errors.New("Method not yet implemented")
}
