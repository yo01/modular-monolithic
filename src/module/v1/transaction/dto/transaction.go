package dto

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	Name      string
	ProductID uuid.UUID
	UserID    uuid.UUID
	CartID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Request
type CreateTransactionRequest struct {
	CartID string `json:"cart_id" validate:"required"`
}

type UpdateTransactionRequest struct {
	CartID string `json:"cart_id" validate:"required"`
}

// Response
type TransactionResponse struct {
	ID               uuid.UUID  `json:"id"`
	CartID           uuid.UUID  `json:"cart_id"`
	IsSuccessPayment bool       `json:"is_success_payment"`
	PaymentDate      *time.Time `json:"payment_date"`
	PaymentByID      string     `json:"payment_by_id"`
	InvoiceNumber    string     `json:"invoice_number"`
}
