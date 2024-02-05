package postgresql

const (
	SELECT_CART = `
		SELECT c.id, c.user_id, c.created_at, cart_item.id as cart_item_id,p.id as product_id, p.name as product_name FROM "cart" c  
			LEFT JOIN cart_item ON cart_item.cart_id = c.id
			LEFT JOIN product p ON cart_item.product_id = p.id
		WHERE c.deleted_at IS NULL
	`

	SELECT_CART_BY_ID = `
		SELECT c.id, c.user_id, cart_item.id as cart_item_id,p.id as product_id, p.name as product_name FROM "cart" c  
			LEFT JOIN cart_item ON cart_item.cart_id = c.id
			LEFT JOIN product p ON cart_item.product_id = p.id
		WHERE c.id = $1 AND c.deleted_at IS NULL
	`

	INSERT_CART = `
		INSERT INTO "cart" 
			("id", "user_id", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
		RETURNING id
	`

	UPDATE_CART = `
		UPDATE "cart"
			SET ("user_id", "updated_at", "updated_by_id") = ($2, NOW(), $3)
		WHERE id = $1 
	`

	HARD_DELETE_CART = `
		DELETE FROM "cart"
		WHERE id = $1;
	`

	SOFT_DELETE_CART = `
		UPDATE "cart"
			SET ("updated_at", "updated_by_id", "deleted_at", "deleted_by_id") = (NOW(), $2, NOW(), $2)
		WHERE id = $1
	`

	SELECT_CART_BY_USER_LOGIN = `
		SELECT * FROM "cart" c
			JOIN cart_item ON cart_item.cart_id = c.id
		WHERE c.user_id = $1 AND c.deleted_at IS NULL
	`

	SELECT_ONE_CART_BY_ID = `
		SELECT c.id, c.user_id, cart_item.id as cart_item_id,p.id as product_id, p.name as product_name FROM "cart" c  
			JOIN cart_item ON cart_item.cart_id = c.id
			JOIN product p ON cart_item.product_id = p.id
		WHERE c.id = $1 AND cart_item.id = $2 AND c.deleted_at IS NULL
	`
)
