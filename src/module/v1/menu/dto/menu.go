package dto

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
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
