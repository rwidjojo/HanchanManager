package service

import (
	"context"
	"fmt"
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

func (s *HanchanService) CreateHanchan(ctx context.Context, groupID int, name *string, date time.Time, uma *[]int, baseScore *int) (*domain.Hanchan, error) {

	var hanchanUma []int
	var hanchanBaseScore int

	if date.IsZero() {
		date = time.Now()
	}

	if uma == nil {
		hanchanUma = []int{15000, 5000, -5000, -15000}
	} else if len(*uma) != 4 {
		return nil, fmt.Errorf("Uma must have exactly 4 values, got %v", uma)
	} else {
		hanchanUma = *uma
	}

	if baseScore == nil {
		hanchanBaseScore = 30000
	} else {
		hanchanBaseScore = *baseScore
	}

	hanchan := &domain.Hanchan{GroupID: groupID, Name: name, Date: date, Uma: hanchanUma, BaseScore: hanchanBaseScore}

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
