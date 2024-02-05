package postgresql

const (
	SELECT_PRODUCT = `
		SELECT p.id, p.name, p.created_at FROM "product" p
			LEFT JOIN menu ON p.menu_id = menu.id
		WHERE p.deleted_at IS NULL
	`

	SELECT_PRODUCT_BY_ID = `
		SELECT * p.id, p.name "product" p WHERE p.id = $1
			LEFT JOIN menu ON p.menu_id = menu.id
		WHERE p.deleted_at IS NULL
	`

	INSERT_PRODUCT = `
		INSERT INTO "product" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_PRODUCT = `
		UPDATE "product"
			SET ("name", "updated_at", "updated_by_id") = ($2, NOW(), $3)
		WHERE id = $1
	`

	HARD_DELETE_PRODUCT = `
		DELETE FROM "product"
		WHERE id = $1;
	`

	SOFT_DELETE_PRODUCT = `
		UPDATE "product"
			SET ("updated_at", "updated_by_id", "deleted_at", "deleted_by_id") = (NOW(), $2, NOW(), $2)
		WHERE id = $1
	`
)
