package postgresql

const (
	SELECT_ROLE = `
		SELECT * FROM "role" u
	`

	SELECT_ROLE_BY_ID = `
		SELECT * FROM "role" u WHERE u.id = $1
	`

	INSERT_ROLE = `
		INSERT INTO "role" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_ROLE = `
		UPDATE "role"
			SET name = $2
		WHERE id = $1
	`

	DELETE_ROLE = `
		DELETE FROM "role"
		WHERE id = $1;
	`
)
