package model

import (
	"time"

	"github.com/google/uuid"
)

// Init Table
type User struct {
	ID        uuid.UUID  `db:"id"`
	Email     string     `db:"email"`
	Password  *string    `db:"password"`
	FullName  string     `db:"full_name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
