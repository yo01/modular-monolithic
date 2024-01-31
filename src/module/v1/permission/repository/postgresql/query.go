package postgresql

const (
	SELECT_PERMISSION = `
		SELECT * FROM "permission" u
	`

	SELECT_PERMISSION_BY_ID = `
		SELECT * FROM "permission" u WHERE u.id = $1
	`

	INSERT_PERMISSION = `
		INSERT INTO "permission" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_PERMISSION = `
		UPDATE "permission"
			SET ("name", "updated_at") = ($2, NOW())
		WHERE id = $1
	`

	DELETE_PERMISSION = `
		DELETE FROM "permission"
		WHERE id = $1;
	`
)
