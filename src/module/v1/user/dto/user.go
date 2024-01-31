package dto

import (
	"modular-monolithic/module/v1/role/dto"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  *string
	FullName  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Request
type CreateUserRequest struct {
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8"`
	FullName string    `json:"full_name" validate:"required"`
	RoleID   uuid.UUID `json:"role_id"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" validate:"required"`
}

// Response
type UserResponse struct {
	ID       uuid.UUID         `json:"id"`
	Email    string            `json:"email"`
	FullName string            `json:"full_name"`
	Role     *dto.RoleResponse `json:"role"`
}
