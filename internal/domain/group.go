package domain

import (
	"time"
)

type Group struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
