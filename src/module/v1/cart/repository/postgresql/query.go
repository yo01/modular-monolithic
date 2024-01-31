package postgresql

const (
	SELECT_CART = `
		SELECT * FROM "cart" u
	`

	SELECT_CART_BY_ID = `
		SELECT * FROM "cart" u WHERE u.id = $1
	`

	INSERT_CART = `
		INSERT INTO "cart" 
			("id", "product_id", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_CART = `
		UPDATE "cart"
			SET ("product_id", "updated_at") = ($2, NOW())	
		WHERE id = $1
	`

	DELETE_CART = `
		DELETE FROM "cart"
		WHERE id = $1;
	`
)
