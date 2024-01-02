package eventDto

import (
	"github.com/google/uuid"
	"time"
)

type EventDto struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	Location    string    `json:"location" binding:"required"`
}

type EventResponseDto struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserId      uuid.UUID `json:"user_id" binding:"required"`
}
