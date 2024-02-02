package dto

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Request
type CreateCartRequest struct {
	ProductID []string `json:"product_id" validate:"required"`
}

type UpdateCartRequest struct {
	ProductID  string `json:"product_id" validate:"required"`
	CartItemID string `json:"cart_item_id" validate:"required"`
}

// Response
type (
	CartResponse struct {
		ID       uuid.UUID           `json:"id"`
		UserID   uuid.UUID           `json:"user_id"`
		CartItem []CartItemReference `json:"cart_item"`
	}

	CartItemReference struct {
		ID          uuid.UUID `json:"id"`
		ProductID   uuid.UUID `json:"product_id"`
		ProductName *string   `json:"product_name"`
	}
)
