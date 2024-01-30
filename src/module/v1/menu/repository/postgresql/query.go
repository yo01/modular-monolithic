package postgresql

const (
	SELECT_MENU = `
		SELECT * FROM "menu" u
	`

	SELECT_MENU_BY_ID = `
		SELECT * FROM "menu" u WHERE u.id = $1
	`

	INSERT_MENU = `
		INSERT INTO "menu" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_MENU = `
		UPDATE "menu"
			SET name = $2
		WHERE id = $1
	`

	DELETE_MENU = `
		DELETE FROM "menu"
		WHERE id = $1;
	`
)
