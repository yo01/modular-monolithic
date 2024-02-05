package postgresql

const (
	SELECT_CART_ITEM = `
		SELECT * FROM "cart_item" c
		WHERE c.deleted_at IS NULL
	`

	SELECT_CART_ITEM_BY_ID = `
		SELECT * FROM "cart_item" c 
		WHERE c.id = $1 AND c.deleted_at IS NULL
	`

	INSERT_CART_ITEM = `
		INSERT INTO "cart_item" 
			("id", "cart_id", "product_id", "created_at", "updated_at", "created_by_id", "updated_by_id")
		VALUES
			($1, $2, $3, NOW(), NOW(), $4, $4)
	`

	UPDATE_CART_ITEM = `
		UPDATE "cart_item"
			SET ("product_id", "updated_at", "updated_by_id") = ($3, NOW(), $4)
		WHERE id = $1 AND cart_id = $2
	`

	HARD_DELETE_CART_ITEM = `
		DELETE FROM "cart_item"
		WHERE id = $1;
	`

	SOFT_DELETE_CART_ITEM = `
		UPDATE "cart_item"
			SET ("updated_at", "updated_by_id", "deleted_at", "deleted_by_id") = (NOW(), $2, NOW(), $2)
		WHERE cart_id = $1
	`
)
