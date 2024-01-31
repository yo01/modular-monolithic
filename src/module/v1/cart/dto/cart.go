package dto

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID
	ProductID string
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
	ProductID string    `json:"product_id"`
}
