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

func (s *HanchanService) CreateHanchan(ctx context.Context, groupID int, name *string, date time.Time, uma *[]int, baseScore *int, seating []domain.PlayerSeating) (*domain.Hanchan, error) {

	var hanchanUma []int
	var hanchanBaseScore int

	if date.IsZero() {
		date = time.Now()
	}

	if uma == nil {
		hanchanUma = []int{15000, 5000, -5000, -15000}
	} else if len(*uma) != 4 {
		return nil, fmt.Errorf("uma must have exactly 4 values, got %v", uma)
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

func (s *HanchanService) GetHanchanByID(ctx context.Context, id int) (*domain.Hanchan, error) {

	hanchan, err := s.hanchanRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return hanchan, nil
}

func (s *HanchanService) ListHanchansByGroupID(ctx context.Context, groupID int) ([]*domain.Hanchan, error) {

	var hanchans []*domain.Hanchan

	hanchans, err := s.hanchanRepo.ListByGroup(ctx, groupID)

	if err != nil {
		return nil, err
	}

	return hanchans, nil
}

func (s *HanchanService) AssignPlayerToHanchan(ctx context.Context, playerID int, hanchanID int, initialSeat domain.SeatWind) error {

	hp := &domain.HanchanPlayer{
		HanchanID:  hanchanID,
		PlayerSeat: domain.PlayerSeating{PlayerID: playerID, InitialSeat: initialSeat},
	}

	return s.hanchanRepo.AssignPlayer(ctx, hp)

}
