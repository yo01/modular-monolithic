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
	Cart             *dto.CartResponse `json:"cart"`
	PaymentDate      *time.Time        `json:"payment_date"`
	ID               uuid.UUID         `json:"id"`
	IsSuccessPayment bool              `json:"is_success_payment"`
}
