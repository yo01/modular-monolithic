package dto

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedAt time.Time
	Name      string
	ID        uuid.UUID
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
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}
