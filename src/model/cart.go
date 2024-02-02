package model

import (
	"time"

	"github.com/google/uuid"
)

// Init Table
type (
	Cart struct {
		ID          uuid.UUID  `db:"id"`
		UserID      uuid.UUID  `db:"user_id"`
		CartItemID  uuid.UUID  `db:"cart_item_id"`
		ProductID   uuid.UUID  `db:"product_id"`
		ProductName *string    `db:"product_name"`
		CreatedAt   time.Time  `db:"created_at"`
		UpdatedAt   *time.Time `db:"updated_at"`
		DeletedAt   *time.Time `db:"deleted_at"`
		CreatedByID uuid.UUID  `db:"created_by_id"`
		UpdatedByID uuid.UUID  `db:"updated_by_id"`
		DeletedByID uuid.UUID  `db:"deleted_by_id"`
	}

	CartItem struct {
		ID          uuid.UUID  `db:"id"`
		CartID      uuid.UUID  `db:"cart_id"`
		ProductID   uuid.UUID  `db:"product_id"`
		CreatedAt   time.Time  `db:"created_at"`
		UpdatedAt   *time.Time `db:"updated_at"`
		DeletedAt   *time.Time `db:"deleted_at"`
		CreatedByID uuid.UUID  `db:"created_by_id"`
		UpdatedByID uuid.UUID  `db:"updated_by_id"`
		DeletedByID uuid.UUID  `db:"deleted_by_id"`
	}
)
