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

type CreateHanchanInput struct {
	GroupID   int
	Name      *string
	Date      time.Time
	BaseScore *int
	Uma       *[]int
	Seating   []domain.PlayerSeating
}

func (s *HanchanService) CreateHanchan(ctx context.Context, input CreateHanchanInput) (*domain.Hanchan, error) {

	var hanchanUma []int
	var hanchanBaseScore int

	hanchanDate := input.Date
	if hanchanDate.IsZero() {
		hanchanDate = time.Now()
	}

	if input.Uma == nil {
		hanchanUma = []int{15000, 5000, -5000, -15000}
	} else if len(*input.Uma) != 4 {
		return nil, fmt.Errorf("uma must have exactly 4 values, got %d", len(*input.Uma))
	} else {
		hanchanUma = *input.Uma
	}

	if input.BaseScore == nil {
		hanchanBaseScore = 30000
	} else {
		hanchanBaseScore = *input.BaseScore
	}

	if err := validateSeating(input.Seating); err != nil {
		return nil, fmt.Errorf("invalid seating: %w", err)
	}

	if err := validatePlayers(input.Seating); err != nil {
		return nil, fmt.Errorf("invalid players: %w", err)
	}

	// ToDo: implement validation to check if requested player id is part of the group
	// if not, need to raise Error

	// ToDo: we should implement transaction here
	// hanchan creation and player assignment should be one single transaction
	hanchan := &domain.Hanchan{
		GroupID:   input.GroupID,
		Name:      input.Name,
		Date:      hanchanDate,
		Uma:       hanchanUma,
		BaseScore: hanchanBaseScore,
	}

	if err := s.hanchanRepo.Create(ctx, hanchan); err != nil {
		return nil, fmt.Errorf("create hanchan: %w", err)
	}

	for _, playerSeat := range input.Seating {
		hp := &domain.HanchanPlayer{HanchanID: hanchan.ID, PlayerSeat: playerSeat}

		if err := s.hanchanRepo.AssignPlayer(ctx, hp); err != nil {
			return nil, fmt.Errorf("assign player: %w", err)
		}

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

func validateSeating(seats []domain.PlayerSeating) error {

	if len(seats) != 4 {
		return fmt.Errorf("player seating must have exactly 4 values, got %d", len(seats))
	}

	seatSeen := make(map[domain.SeatWind]int, 4)

	for _, seat := range seats {
		switch seat.InitialSeat {
		case domain.SeatEast, domain.SeatSouth, domain.SeatWest, domain.SeatNorth:
			seatSeen[seat.InitialSeat]++
		default:
			return fmt.Errorf("invalid seat: %s", seat.InitialSeat)
		}
	}

	for _, s := range []domain.SeatWind{domain.SeatEast, domain.SeatSouth, domain.SeatWest, domain.SeatNorth} {
		if seatSeen[s] != 1 {
			return fmt.Errorf("seat %s must appear exactly once", s)
		}
	}

	return nil

}

func validatePlayers(seats []domain.PlayerSeating) error {

	playerSeen := make(map[int]struct{}, 4)

	for _, seat := range seats {
		if _, exists := playerSeen[seat.PlayerID]; exists {
			return fmt.Errorf("found duplicate player_id: %d", seat.PlayerID)
		}
		playerSeen[seat.PlayerID] = struct{}{}
	}

	return nil

}
