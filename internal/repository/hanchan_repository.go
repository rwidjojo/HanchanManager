package repository

import (
	"context"

	"HanchanManager/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HanchanRepository interface {
	Create(ctx context.Context, hanchan *domain.Hanchan) error
	GetByID(ctx context.Context, id int) (*domain.Hanchan, error)
	ListByGroup(ctx context.Context, groupID int) ([]domain.Hanchan, error)
	AddPlayer(ctx context.Context, hp *domain.HanchanPlayer) error
	ListPlayers(ctx context.Context, hanchanID int) ([]domain.HanchanPlayer, error)
	Close(ctx context.Context, hanchanID int) error
}

type hanchanRepo struct {
	db *pgxpool.Pool
}

func NewHanchanRepo(db *pgxpool.Pool) HanchanRepository {
	return &hanchanRepo{db: db}
}
