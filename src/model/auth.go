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
	ID       uuid.UUID
	FullName string
	Role     dto.RoleResponse
}

type Claims struct {
	Authorized bool      `json:"autorized"`
	UserID     uuid.UUID `json:"user_id"`
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	RoleID     uuid.UUID `json:"role_id"`
	RoleName   string    `json:"role_name"`
}
