package dto

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedAt time.Time
	ID        uuid.UUID
	ProductID uuid.UUID
	UserID    uuid.UUID
}

// Request
type CreateCartItemRequest struct {
	CartID    string `json:"cart_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
}

type UpdateCartItemRequest struct {
	ProductID string `json:"product_id" validate:"required"`
}

// Response
type CartItemResponse struct {
	CartID    uuid.UUID `json:"cart_id"`
	ProductID uuid.UUID `json:"product_id"`
	ID        uuid.UUID `json:"id"`
}
