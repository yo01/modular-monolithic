package dto

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedAt time.Time
	Name      string
	UserID    []uuid.UUID
	ID        uuid.UUID
}

// Request
type CreatePermissionRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePermissionRequest struct {
	Name string `json:"name" validate:"required"`
}

// Response
type PermissionResponse struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}
