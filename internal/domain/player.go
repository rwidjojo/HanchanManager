package domain

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
