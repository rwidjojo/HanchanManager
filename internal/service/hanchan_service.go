package service

import (
	"context"
	"time"

	"HanchanManager/internal/domain"
	"HanchanManager/internal/repository"
)

type HanchanService struct {
	hanchanRepo repository.HanchanRepository
}

func NewHanchanService(repo repository.HanchanRepository) *HanchanService {
	return &HanchanService{hanchanRepo: repo}
}

func (s *HanchanService) CreateHanchan(ctx context.Context, group_id int, name *string, date time.Time, uma []int) (*domain.Hanchan, error) {

	hanchan := &domain.Hanchan{GroupID: group_id, Name: name, Date: date, Uma: uma}

	if err := s.hanchanRepo.Create(ctx, hanchan); err != nil {
		return nil, err
	}

	return hanchan, nil
}

func (s *HanchanService) GetByID(ctx context.Context, id int) (*domain.Hanchan, error) {

	hanchan, err := s.hanchanRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return hanchan, nil
}
