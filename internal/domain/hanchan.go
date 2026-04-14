package domain

import (
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
	HanchanID   int      `json:"hanchan_id"`
	PlayerID    int      `json:"player_id"`
	InitialSeat SeatWind `json:"initial_seat"`
	FinalScore  *int     `json:"final_score,omitempty"` // null until hanchan closed
	Placement   *int     `json:"placement,omitempty"`   // null until hanchan closed
}
