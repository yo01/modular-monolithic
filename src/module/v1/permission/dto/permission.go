package dto

import (
	roleDTO "modular-monolithic/module/v1/role/dto"
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
	RoleID  string   `json:"role_id"`
	ListAPI []string `json:"list_api"`
}

type UpdatePermissionRequest struct {
	RoleID  string   `json:"role_id"`
	ListAPI []string `json:"list_api"`
}

// Response
type PermissionResponse struct {
	Role    *roleDTO.RoleResponse `json:"role,omitempty"`
	ListAPI []string              `json:"list_api"`
	ID      uuid.UUID             `json:"id"`
}
