package model

import (
	"time"

	"github.com/google/uuid"
)

// Init Table
type Permission struct {
	ID        uuid.UUID  `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
