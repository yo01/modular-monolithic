package postgresql

const (
	SELECT_USER = `
		SELECT * FROM "user" u
	`

	SELECT_USER_BY_ID = `
		SELECT * FROM "user" u WHERE u.id = $1
	`

	INSERT_USER = `
		INSERT INTO "user" 
			("id", "email", "password", "full_name", "created_at", "updated_at")
		VALUES
			($1, $2, $3, $4, NOW(), NOW())
	`

	UPDATE_USER = `
		UPDATE "user"
			SET ("full_name", "updated_at") = ($2, NOW())
		WHERE id = $1
	`

	DELETE_USER = `
		DELETE FROM "user"
		WHERE id = $1;
	`

	// ADDITIONAL
	SELECT_USER_BY_EMAIL = `
		SELECT * FROM "user" u WHERE u.email = $1
	`
)
