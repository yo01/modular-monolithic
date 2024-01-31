package postgresql

const (
	SELECT_TRANSACTION = `
		SELECT * FROM "transaction" u
	`

	SELECT_TRANSACTION_BY_ID = `
		SELECT * FROM "transaction" u WHERE u.id = $1
	`

	INSERT_TRANSACTION = `
		INSERT INTO "transaction" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_TRANSACTION = `
		UPDATE "transaction"
			SET ("name", "updated_at") = ($2, NOW())
		WHERE id = $1
	`

	DELETE_TRANSACTION = `
		DELETE FROM "transaction"
		WHERE id = $1;
	`

	// ADDITIONAL
	UPDATE_TRANSACTION_PAYMENT = `
		UPDATE "transaction"
			SET ("is_success_payment", "updated_at") = ($2, NOW())
		WHERE id = $1
	`
)
