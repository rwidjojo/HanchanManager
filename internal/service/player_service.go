package service

import (
	"context"
	"errors"

	"HanchanManager/internal/domain"
	"HanchanManager/internal/repository"
)

type PlayerService struct {
	repo repository.PlayerRepository
}

func NewPlayerService(repo repository.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) CreatePlayer(ctx context.Context, username string, name string) (*domain.Player, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}

	if name == "" {
		return nil, errors.New("name is required")
	}

	player := &domain.Player{Username: username, Name: name}

	if err := s.repo.Create(ctx, player); err != nil {
		return nil, err
	}

	return player, nil
}

func (s *PlayerService) ListPlayers(ctx context.Context) ([]*domain.Player, error) {

	var players []*domain.Player

	players, err := s.repo.List(ctx)

	if err != nil {
		return nil, err
	}

	return players, nil
}

func (s *PlayerService) GetPlayerByID(ctx context.Context, id int) (*domain.Player, error) {

	player, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return player, nil
}
