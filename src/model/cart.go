package model

import (
	"time"

	"github.com/google/uuid"
)

// Init Table
type Cart struct {
	ID        uuid.UUID  `db:"id"`
	ProductID string     `db:"product_id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
