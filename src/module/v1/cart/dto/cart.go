package dto

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Request
type CreateCartRequest struct {
	ProductID string `json:"product_id" validate:"required"`
}

type UpdateCartRequest struct {
	ProductID string `json:"product_id" validate:"required"`
}

// Response
type CartResponse struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	UserID    uuid.UUID `json:"user_id"`
}
