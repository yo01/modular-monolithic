package postgresql

const (
	SELECT_MENU = `
		SELECT m.id, m.name, m.created_at, m.deleted_at FROM "menu" m
		WHERE m.deleted_at IS NULL
	`

	SELECT_MENU_BY_ID = `
		SELECT m.id, m.name, m.created_at, m.deleted_at FROM "menu" m
		WHERE m.deleted_at IS NULL AND m.id = $1
	`

	INSERT_MENU = `
		INSERT INTO "menu" 
			("id", "name", "created_at", "updated_at", "created_by_id", "updated_by_id")
		VALUES
			($1, $2, NOW(), NOW(), $3, $3)
	`

	UPDATE_MENU = `
		UPDATE "menu"
			SET ("name", "updated_at", "updated_by_id") = ($2, NOW(), $3)
		WHERE id = $1
	`

	HARD_DELETE_MENU = `
		DELETE FROM "menu"
		WHERE id = $1;
	`

	SOFT_DELETE_MENU = `
		UPDATE "menu"
			SET ("updated_at", "updated_by_id", "deleted_at", "deleted_by_id") = (NOW(), $2, NOW(), $2)
		WHERE id = $1
	`
)
