package postgresql

const (
	SELECT_PRODUCT = `
		SELECT * FROM "product" p
		WHERE p.deleted_at IS NULL
	`

	SELECT_PRODUCT_BY_ID = `
		SELECT * FROM "product" p WHERE p.id = $1
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
			SET ("name", "updated_at", "updated_by_id", "updated_by_full_name") = ($2, NOW(), $3, $4)
		WHERE id = $1
	`

	HARD_DELETE_PRODUCT = `
		DELETE FROM "product"
		WHERE id = $1;
	`

	SOFT_DELETE_PRODUCT = `
		UPDATE "product"
			SET ("updated_at", "updated_by_id", "updated_by_full_name", "deleted_at", "deleted_by_id", "deleted_by_full_name") = (NOW(), $2, $3, NOW(), $2, $3)
		WHERE id = $1
	`
)
