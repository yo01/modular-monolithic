package dto

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Request
type CreateProductRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateProductRequest struct {
	Name string `json:"name" validate:"required"`
}

// Response
type ProductResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
