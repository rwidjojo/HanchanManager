package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

type HanchanStatus string

const (
	HanchanOpen   HanchanStatus = "OPEN"
	HanchanClosed HanchanStatus = "CLOSED"
)

type SeatWind string

const (
	SeatEast  SeatWind = "EAST"
	SeatSouth SeatWind = "SOUTH"
	SeatWest  SeatWind = "WEST"
	SeatNorth SeatWind = "NORTH"
)

func (sw SeatWind) IsValid() bool {
	switch sw {
	case SeatEast, SeatSouth, SeatWest, SeatNorth:
		return true
	default:
		return false
	}
}

func (sw *SeatWind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch SeatWind(s) {
	case SeatEast, SeatSouth, SeatWest, SeatNorth:
		*sw = SeatWind(s)
		return nil
	default:
		return fmt.Errorf("invalid SeatWind: %s (allowed: EAST, SOUTH, WEST, NORTH)", s)
	}
}

type PlayerSeating struct {
	PlayerID    int      `json:"player_id"`
	InitialSeat SeatWind `json:"initial_seat"`
}

type Hanchan struct {
	ID        int            `json:"id"`
	GroupID   int            `json:"group_id"`
	Name      *string        `json:"name,omitempty"`
	Date      time.Time      `json:"date"`
	Status    *HanchanStatus `json:"status,omitempty"`
	BaseScore int            `json:"base_score,omitempty"`
	Uma       []int          `json:"uma"` // [1st, 2nd, 3rd, 4th] point adjustments
	CreatedAt time.Time      `json:"created_at"`
}

type HanchanPlayer struct {
	HanchanID  int           `json:"hanchan_id"`
	PlayerSeat PlayerSeating `json:"player_seat"`
	FinalScore *int          `json:"final_score,omitempty"` // null until hanchan closed
	Placement  *int          `json:"placement,omitempty"`   // null until hanchan closed
}
