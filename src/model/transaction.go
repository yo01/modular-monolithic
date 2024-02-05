package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Init Table
type Transaction struct {
	UpdatedAt        *time.Time   `db:"updated_at"`
	DeletedAt        *time.Time   `db:"deleted_at"`
	PaymentDate      *time.Time   `db:"payment_date"`
	CreatedAt        time.Time    `db:"created_at"`
	IsSuccessPayment sql.NullBool `db:"is_success_payment"`
	PaymentByID      string       `db:"payment_by_id"`
	InvoiceNumber    string       `db:"invoice_number"`
	ID               uuid.UUID    `db:"id"`
	CartID           uuid.UUID    `db:"cart_id"`
	CartUserID       uuid.UUID    `db:"cart_user_id"`
	CreatedByID      uuid.UUID    `db:"created_by_id"`
	UpdatedByID      uuid.UUID    `db:"updated_by_id"`
	DeletedByID      uuid.UUID    `db:"deleted_by_id"`
}
