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
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_TRANSACTION = `
		UPDATE "transaction"
			SET ("name", "updated_at", "updated_by_id", "updated_by_full_name") = ($2, NOW(), $3, $4)
		WHERE id = $1
	`

	HARD_DELETE_TRANSACTION = `
		DELETE FROM "transaction"
		WHERE id = $1;
	`

	SOFT_DELETE_TRANSACTION = `
		UPDATE "transaction"
			SET ("updated_at", "updated_by_id", "updated_by_full_name", "deleted_at", "deleted_by_id", "deleted_by_full_name") = (NOW(), $2, $3, NOW(), $2, $3)
		WHERE id = $1
	`

	// ADDITIONAL
	UPDATE_TRANSACTION_PAYMENT = `
		UPDATE "transaction"
			SET ("is_success_payment", "payment_date", "payment_by_id", "payment_by_full_name", "updated_at", "updated_by_id", "updated_by_full_name") = ($2, $3, $4, NOW(), NOW(), $3, $4)
		WHERE id = $1
	`
)
