package postgresql

const (
	SELECT_USER = `
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, role.name as role_name  FROM "user" u
			LEFT JOIN role ON u.role_id = role.id 
	`

	SELECT_USER_BY_ID = `
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, role.name as role_name  FROM "user" u
			LEFT JOIN role ON u.role_id = role.id 
		WHERE u.id = $1
	`

	INSERT_USER = `
		INSERT INTO "user" 
			("id", "email", "password", "full_name", "role_id", "created_at", "updated_at")
		VALUES
			($1, $2, $3, $4, $5, NOW(), NOW())
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
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, role.name as role_name  FROM "user" u
			LEFT JOIN role ON u.role_id = role.id 
		WHERE u.email = $1
	`
)
