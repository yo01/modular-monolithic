package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Init Table
type User struct {
	ListAPI          pq.StringArray `db:"list_api"`
	UpdatedAt        *time.Time     `db:"updated_at"`
	DeletedAt        *time.Time     `db:"deleted_at"`
	CreatedAt        time.Time      `db:"created_at"`
	Password         *string        `db:"password"`
	RoleName         *string        `db:"role_name"`
	Email            string         `db:"email"`
	FullName         string         `db:"full_name"`
	ID               uuid.UUID      `db:"id"`
	PermissionID     uuid.UUID      `db:"permission_id"`
	PermissionRoleID uuid.UUID      `db:"permission_role_id"`
	CreatedByID      uuid.UUID      `db:"created_by_id"`
	UpdatedByID      uuid.UUID      `db:"updated_by_id"`
	DeletedByID      uuid.UUID      `db:"deleted_by_id"`
	RoleID           uuid.UUID      `db:"role_id"`
}
