package dto

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Name      string
	ID        uuid.UUID
}

// Request
type CreateMenuRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateMenuRequest struct {
	Name string `json:"name" validate:"required"`
}

// Response
type MenuResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
