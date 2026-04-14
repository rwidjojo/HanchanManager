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

func (s *HanchanService) CreateHanchan(ctx context.Context, group_id int, name *string, date time.Time, uma *[]int, base_score *int) (*domain.Hanchan, error) {

	var hanchan_uma []int
	var hanchan_base_score int

	if date.IsZero() {
		date = time.Now()
	}

	if uma == nil {
		hanchan_uma = []int{15000, 5000, -5000, -15000}
	} else if len(*uma) != 4 {
		return nil, fmt.Errorf("Uma must have exactly 4 values, got %v", uma)
	} else {
		hanchan_uma = *uma
	}

	if base_score == nil {
		hanchan_base_score = 30000
	}

	hanchan := &domain.Hanchan{GroupID: group_id, Name: name, Date: date, Uma: hanchan_uma, BaseScore: hanchan_base_score}

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
