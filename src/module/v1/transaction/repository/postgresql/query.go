package postgresql

const (
	SELECT_TRANSACTION = `
		SELECT t.id, t.is_success_payment, t.payment_date, t.created_at, t.deleted_at, cart.id as cart_id, cart.user_id as cart_user_id FROM "transaction" t
			LEFT JOIN cart ON t.cart_id = cart.id 
		WHERE t.deleted_at IS NULL
	`

	SELECT_TRANSACTION_BY_ID = `
		SELECT t.id, t.is_success_payment, t.payment_date, t.created_at, t.deleted_at, cart.id as cart_id, cart.user_id as cart_user_id FROM "transaction" t 
			LEFT JOIN cart ON t.cart_id = cart.id 
		WHERE t.id = $1 AND t.deleted_at IS NULL
	`

	INSERT_TRANSACTION = `
		INSERT INTO "transaction" 
			("id", "cart_id", "is_success_payment", "payment_date", "created_at", "updated_at", "created_by_id", "updated_by_id")
		VALUES
			($1, $2, $3, NOW(), NOW(), NOW(), $4, $4)
	`

	UPDATE_TRANSACTION = `
		UPDATE "transaction"
			SET ("cart_id", "updated_at", "updated_by_id") = ($2, NOW(), $3)
		WHERE id = $1
	`

	HARD_DELETE_TRANSACTION = `
		DELETE FROM "transaction"
		WHERE id = $1;
	`

	SOFT_DELETE_TRANSACTION = `
		UPDATE "transaction"
			SET ("updated_at", "updated_by_id", "deleted_at", "deleted_by_id") = (NOW(), $2, NOW(), $2)
		WHERE id = $1
	`

	// ADDITIONAL
	UPDATE_TRANSACTION_PAYMENT = `
		UPDATE "transaction"
			SET ("is_success_payment", "payment_date", "invoice_number", "updated_at", "updated_by_id") = ($2, $3, NOW(), NOW(), $4)
		WHERE id = $1
	`
)
