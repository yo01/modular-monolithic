package postgresql

const (
	SELECT_MENU = `
		SELECT * FROM "menu" m
		WHERE m.deleted_at IS NULL
	`

	SELECT_MENU_BY_ID = `
		SELECT * FROM "menu" m WHERE m.id = $1
		WHERE m.deleted_at IS NULL
	`

	INSERT_MENU = `
		INSERT INTO "menu" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_MENU = `
		UPDATE "menu"
			SET ("name", "updated_at", "updated_by_id", "updated_by_full_name") = ($2, NOW(), $3, $4)
		WHERE id = $1
	`

	HARD_DELETE_MENU = `
		DELETE FROM "menu"
		WHERE id = $1;
	`

	SOFT_DELETE_MENU = `
		UPDATE "menu"
			SET ("updated_at", "updated_by_id", "updated_by_full_name", "deleted_at", "deleted_by_id", "deleted_by_full_name") = (NOW(), $2, $3, NOW(), $2, $3)
		WHERE id = $1
	`
)
