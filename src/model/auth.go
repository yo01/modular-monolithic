package model

import (
	"modular-monolithic/module/v1/role/dto"

	"github.com/google/uuid"
)

type Auth struct {
	User  AuthUser
	Token string
}

type AuthUser struct {
	Role     *dto.RoleResponse
	FullName string
	ID       uuid.UUID
}

type Claims struct {
	Role       *dto.RoleResponse `json:"role"`
	Email      string            `json:"email"`
	FullName   string            `json:"full_name"`
	UserID     uuid.UUID         `json:"user_id"`
	Authorized bool              `json:"autorized"`
}
