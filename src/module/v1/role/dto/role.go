package dto

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
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
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
