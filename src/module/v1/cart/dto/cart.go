package dto

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedAt time.Time
	ID        uuid.UUID
	ProductID uuid.UUID
	UserID    uuid.UUID
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
		CartItem []CartItemReference `json:"cart_item"`
		ID       uuid.UUID           `json:"id"`
		UserID   uuid.UUID           `json:"user_id"`
	}

	CartItemReference struct {
		ProductName *string   `json:"product_name"`
		ID          uuid.UUID `json:"id"`
		ProductID   uuid.UUID `json:"product_id"`
	}
)
