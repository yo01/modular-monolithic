package dto

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedAt time.Time
	Name      string
	ID        uuid.UUID
}

// Request
type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}

// Response
type RoleResponse struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}
