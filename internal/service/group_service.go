package service

import (
	"context"
	"errors"

	"github.com/rwidjojo/HanchanManager/internal/domain"
	"github.com/rwidjojo/HanchanManager/internal/repository"
)

type GroupService struct {
	repo repository.GroupRepository
}

func NewGroupService(repo repository.GroupRepository) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) CreateGroup(ctx context.Context, code string, description *string) (*domain.Group, error) {
	if code == "" {
		return nil, errors.New("Group unique code is required!")
	}

	group := &domain.Group{Code: code, Description: description}

	if err := s.repo.Create(ctx, group); err != nil {
		return nil, err
	}

	return group, nil
}

func (s *GroupService) GetGroupByCode(ctx context.Context, code string) (*domain.Group, error) {

	group, err := s.repo.GetByCode(ctx, code)

	if err != nil {
		return nil, err
	}

	return group, nil
}
