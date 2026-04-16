package service

import (
	"context"
	"errors"

	"HanchanManager/internal/domain"
	"HanchanManager/internal/repository"
)

type GroupService struct {
	groupRepo      repository.GroupRepository
	membershipRepo repository.MembershipRepository
}

func NewGroupService(groupRepo repository.GroupRepository, membershipRepo repository.MembershipRepository) *GroupService {
	return &GroupService{groupRepo: groupRepo, membershipRepo: membershipRepo}
}

func (s *GroupService) CreateGroup(ctx context.Context, code string, description *string) (*domain.Group, error) {
	if code == "" {
		return nil, errors.New("group unique code is required!")
	}

	group := &domain.Group{Code: code, Description: description}

	if err := s.groupRepo.Create(ctx, group); err != nil {
		return nil, err
	}

	return group, nil
}

func (s *GroupService) GetGroupByID(ctx context.Context, id int) (*domain.Group, error) {

	group, err := s.groupRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return group, nil
}

func (s *GroupService) AddPlayer(ctx context.Context, groupID int, playerID int) error {
	return s.membershipRepo.AddPlayer(ctx, groupID, playerID)
}

func (s *GroupService) GetPlayers(ctx context.Context, groupID int) ([]*domain.Player, error) {
	return s.membershipRepo.GetPlayers(ctx, groupID)
}
