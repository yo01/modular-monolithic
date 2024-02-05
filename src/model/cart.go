package model

import (
	"time"

	"github.com/google/uuid"
)

// Init Table
type (
	Cart struct {
		UpdatedAt   *time.Time `db:"updated_at"`
		DeletedAt   *time.Time `db:"deleted_at"`
		CreatedAt   time.Time  `db:"created_at"`
		ProductName *string    `db:"product_name"`
		ID          uuid.UUID  `db:"id"`
		UserID      uuid.UUID  `db:"user_id"`
		CartItemID  uuid.UUID  `db:"cart_item_id"`
		ProductID   uuid.UUID  `db:"product_id"`
		CreatedByID uuid.UUID  `db:"created_by_id"`
		UpdatedByID uuid.UUID  `db:"updated_by_id"`
		DeletedByID uuid.UUID  `db:"deleted_by_id"`
	}

	CartItem struct {
		UpdatedAt   *time.Time `db:"updated_at"`
		DeletedAt   *time.Time `db:"deleted_at"`
		CreatedAt   time.Time  `db:"created_at"`
		ID          uuid.UUID  `db:"id"`
		CartID      uuid.UUID  `db:"cart_id"`
		ProductID   uuid.UUID  `db:"product_id"`
		CreatedByID uuid.UUID  `db:"created_by_id"`
		UpdatedByID uuid.UUID  `db:"updated_by_id"`
		DeletedByID uuid.UUID  `db:"deleted_by_id"`
	}
)
