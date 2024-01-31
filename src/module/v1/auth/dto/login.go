package dto

import (
	"modular-monolithic/module/v1/role/dto"

	"time"

	"github.com/google/uuid"
)

type Login struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Request
type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Response
type LoginResponse struct {
	ID       uuid.UUID         `json:"id"`
	FullName string            `json:"full_name"`
	Email    string            `json:"email"`
	Role     *dto.RoleResponse `json:"role"`
	Token    string            `json:"token"`
}
