package postgresql

const (
	SELECT_CART = `
		SELECT * FROM "cart" c
		WHERE c.deleted_at IS NULL
	`

	SELECT_CART_BY_ID = `
		SELECT * FROM "cart" c WHERE c.id = $1
		WHERE c.deleted_at IS NULL
	`

	INSERT_CART = `
		INSERT INTO "cart" 
			("id", "product_id", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_CART = `
		UPDATE "cart"
			SET ("product_id", "updated_at", "updated_by_id", "updated_by_full_name") = ($2, NOW(), $3, $4)
		WHERE id = $1
	`

	HARD_DELETE_CART = `
		DELETE FROM "cart"
		WHERE id = $1;
	`

	SOFT_DELETE_CART = `
		UPDATE "cart"
			SET ("updated_at", "updated_by_id", "updated_by_full_name", "deleted_at", "deleted_by_id", "deleted_by_full_name") = (NOW(), $2, $3, NOW(), $2, $3)
		WHERE id = $1
	`
)
