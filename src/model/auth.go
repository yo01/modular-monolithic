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
	Role     dto.RoleResponse
	FullName string
	ID       uuid.UUID
}

type Claims struct {
	UserID     uuid.UUID `json:"user_id"`
	RoleID     uuid.UUID `json:"role_id"`
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	RoleName   string    `json:"role_name"`
	Authorized bool      `json:"autorized"`
}
