package postgresql

const (
	SELECT_TRANSACTION = `
		SELECT * FROM "transaction" t
		WHERE t.deleted_at IS NULL
	`

	SELECT_TRANSACTION_BY_ID = `
		SELECT * FROM "transaction" t WHERE t.id = $1
		WHERE t.deleted_at IS NULL
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
