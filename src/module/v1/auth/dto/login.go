package dto

import (
	"modular-monolithic/module/v1/role/dto"

	"time"

	"github.com/google/uuid"
)

type Login struct {
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedAt time.Time
	Email     string
	Password  string
	ID        uuid.UUID
}

// Request
type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Response
type LoginResponse struct {
	Role         *dto.RoleResponse `json:"role"`
	FullName     string            `json:"full_name"`
	Email        string            `json:"email"`
	AccessToken  string            `json:"access_token"`
	RefreshToken string            `json:"refresh_token"`
	ID           uuid.UUID         `json:"id"`
}
