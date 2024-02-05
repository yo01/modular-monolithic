package dto

import (
	"modular-monolithic/module/v1/cart/dto"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	UpdatedAt *time.Time
	DeletedAt *time.Time
	CreatedAt time.Time
	Name      string
	ID        uuid.UUID
	ProductID uuid.UUID
	UserID    uuid.UUID
	CartID    uuid.UUID
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
	PaymentDate      *time.Time        `json:"payment_date"`
	PaymentByID      string            `json:"payment_by_id"`
	InvoiceNumber    string            `json:"invoice_number"`
	ID               uuid.UUID         `json:"id"`
	Cart             *dto.CartResponse `json:"cart"`
	IsSuccessPayment bool              `json:"is_success_payment"`
}
