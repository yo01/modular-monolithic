package dto

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Request
type CreateTransactionRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateTransactionRequest struct {
	Name string `json:"name" validate:"required"`
}

// Response
type TransactionResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	IsSuccessPayment bool      `json:"is_success_payment"`
}
