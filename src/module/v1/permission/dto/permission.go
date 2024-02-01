package dto

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID
	Name      string
	UserID    []uuid.UUID
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
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
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
