package service

import (
	"context"
	"errors"

	"github.com/rwidjojo/HanchanManager/internal/domain"
	"github.com/rwidjojo/HanchanManager/internal/repository"
)

type PlayerService struct {
	repo repository.PlayerRepository
}

func NewPlayerService(repo repository.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) CreatePlayer(ctx context.Context, username string, name string) (*domain.Player, error) {
	if username == "" {
		return nil, errors.New("Username is required")
	}

	if name == "" {
		return nil, errors.New("Name is required")
	}

	player := &domain.Player{Username: username, Name: name}

	if err := s.repo.Create(ctx, player); err != nil {
		return nil, err
	}

	return player, nil
}
