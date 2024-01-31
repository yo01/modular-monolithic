package postgresql

const (
	SELECT_PRODUCT = `
		SELECT * FROM "product" u
	`

	SELECT_PRODUCT_BY_ID = `
		SELECT * FROM "product" u WHERE u.id = $1
	`

	INSERT_PRODUCT = `
		INSERT INTO "product" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_PRODUCT = `
		UPDATE "product"
			SET ("name", "updated_at") = ($2, NOW())
		WHERE id = $1
	`

	DELETE_PRODUCT = `
		DELETE FROM "product"
		WHERE id = $1;
	`
)
