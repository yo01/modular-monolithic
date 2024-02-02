package model

import (
	"time"

	"github.com/google/uuid"
)

// Init Table
type Product struct {
	ID          uuid.UUID  `db:"id"`
	Name        string     `db:"name"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
	CreatedByID uuid.UUID  `db:"created_by_id"`
	UpdatedByID uuid.UUID  `db:"updated_by_id"`
	DeletedByID uuid.UUID  `db:"deleted_by_id"`
	MenuID      uuid.UUID  `db:"menu_id"`
}
