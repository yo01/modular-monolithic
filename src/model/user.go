package model

import (
	"time"

	"github.com/google/uuid"
)

// Init Table
type User struct {
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
	CreatedAt   time.Time  `db:"created_at"`
	Password    *string    `db:"password"`
	RoleName    *string    `db:"role_name"`
	Email       string     `db:"email"`
	FullName    string     `db:"full_name"`
	ID          uuid.UUID  `db:"id"`
	CreatedByID uuid.UUID  `db:"created_by_id"`
	UpdatedByID uuid.UUID  `db:"updated_by_id"`
	DeletedByID uuid.UUID  `db:"deleted_by_id"`
	RoleID      uuid.UUID  `db:"role_id"`
}
