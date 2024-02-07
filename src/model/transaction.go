package model

import (
	"time"

	"github.com/google/uuid"
)

// Init Table
type Transaction struct {
	UpdatedAt         *time.Time `db:"updated_at"`
	DeletedAt         *time.Time `db:"deleted_at"`
	PaymentDate       *time.Time `db:"payment_date"`
	CreatedAt         time.Time  `db:"created_at"`
	IsSuccessPayment  bool       `db:"is_success_payment"`
	CartIsSuccess     bool       `db:"cart_is_success"`
	PaymentByID       string     `db:"payment_by_id"`
	ProductName       string     `db:"product_name"`
	ID                uuid.UUID  `db:"id"`
	CartID            uuid.UUID  `db:"cart_id"`
	CartUserID        uuid.UUID  `db:"cart_user_id"`
	CartItemID        uuid.UUID  `db:"cart_item_id"`
	CartItemProductID uuid.UUID  `db:"cart_item_product_id"`
	CreatedByID       uuid.UUID  `db:"created_by_id"`
	UpdatedByID       uuid.UUID  `db:"updated_by_id"`
	DeletedByID       uuid.UUID  `db:"deleted_by_id"`
}
