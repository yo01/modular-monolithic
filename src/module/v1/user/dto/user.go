package dto

import (
	"modular-monolithic/module/v1/role/dto"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedAt time.Time
	Password  *string
	Email     string
	FullName  string
	ID        uuid.UUID
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
	Role     *dto.RoleResponse `json:"role"`
	Email    string            `json:"email"`
	FullName string            `json:"full_name"`
	ID       uuid.UUID         `json:"id"`
}

type UserLoginResponse struct {
	Role     *dto.RoleResponse `json:"role"`
	Email    string            `json:"email"`
	FullName string            `json:"full_name"`
	Password string            `json:"password"`
	ID       uuid.UUID         `json:"id"`
}
