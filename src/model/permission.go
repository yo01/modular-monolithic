package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Init Table
type Permission struct {
	ListAPI     pq.StringArray `db:"list_api"`
	UpdatedAt   *time.Time     `db:"updated_at"`
	DeletedAt   *time.Time     `db:"deleted_at"`
	CreatedAt   time.Time      `db:"created_at"`
	RoleName    *string        `db:"role_name"`
	ID          uuid.UUID      `db:"id"`
	RoleID      uuid.UUID      `db:"role_id"`
	CreatedByID uuid.UUID      `db:"created_by_id"`
	UpdatedByID uuid.UUID      `db:"updated_by_id"`
	DeletedByID uuid.UUID      `db:"deleted_by_id"`
}
