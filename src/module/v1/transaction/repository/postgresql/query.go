package postgresql

const (
	SELECT_TRANSACTION_ADMIN = `
		SELECT t.id, t.is_success_payment, t.payment_date, t.created_at, t.deleted_at, cart.id as cart_id, cart.user_id as cart_user_id, cart.is_success as cart_is_success, cart_item.id as cart_item_id, cart_item.product_id as cart_item_product_id, product.name as product_name FROM "transaction" t
			LEFT JOIN cart ON t.cart_id = cart.id
			LEFT JOIN cart_item ON t.cart_id = cart_item.cart_id
			LEFT JOIN product ON cart_item.product_id = product.id
		WHERE t.deleted_at IS NULL
		ORDER BY t.created_at desc LIMIT 10 OFFSET 0
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
	SELECT_TRANSACTION_BY_ID = `
		SELECT t.id, t.is_success_payment, t.payment_date, t.created_at, t.deleted_at, cart.id as cart_id, cart.user_id as cart_user_id,  cart.is_success as cart_is_success, cart_item.id as cart_item_id, cart_item.product_id as cart_item_product_id, product.name as product_name FROM "transaction" t
			LEFT JOIN cart ON t.cart_id = cart.id
			LEFT JOIN cart_item ON t.cart_id = cart_item.cart_id
			LEFT JOIN product ON cart_item.product_id = product.id
		WHERE t.deleted_at IS NULL AND t.id = $1
		ORDER BY t.created_at desc LIMIT 10 OFFSET 0
	`

	UPDATE_TRANSACTION_PAYMENT = `
		UPDATE "transaction"
			SET ("is_success_payment", "payment_date", "invoice_number", "updated_at", "updated_by_id") = ($2, $3, NOW(), NOW(), $4)
		WHERE id = $1
	`

	SELECT_TRANSACTION_LEARNER = `
		SELECT t.id, t.is_success_payment, t.payment_date, t.created_at, t.deleted_at, cart.id as cart_id, cart.user_id as cart_user_id,  cart.is_success as cart_is_success, cart_item.id as cart_item_id, cart_item.product_id as cart_item_product_id, product.name as product_name FROM "transaction" t
			LEFT JOIN cart ON t.cart_id = cart.id
			LEFT JOIN cart_item ON t.cart_id = cart_item.cart_id
			LEFT JOIN product ON cart_item.product_id = product.id
		WHERE t.deleted_at IS NULL AND cart.user_id = $1
		ORDER BY t.created_at desc LIMIT 10 OFFSET 0
	`
)
