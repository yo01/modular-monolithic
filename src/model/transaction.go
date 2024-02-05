package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Init Table
type Transaction struct {
	ID               uuid.UUID    `db:"id"`
	CartID           uuid.UUID    `db:"cart_id"`
	PaymentByID      string       `db:"payment_by_id"`
	IsSuccessPayment sql.NullBool `db:"is_success_payment"`
	PaymentDate      *time.Time   `db:"payment_date"`
	InvoiceNumber    string       `db:"invoice_number"`
	CreatedAt        time.Time    `db:"created_at"`
	UpdatedAt        *time.Time   `db:"updated_at"`
	DeletedAt        *time.Time   `db:"deleted_at"`
	CreatedByID      uuid.UUID    `db:"created_by_id"`
	UpdatedByID      uuid.UUID    `db:"updated_by_id"`
	DeletedByID      uuid.UUID    `db:"deleted_by_id"`
}
