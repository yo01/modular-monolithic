package postgresql

const (
	SELECT_USER = `
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, u.created_at, u.created_at, u.deleted_at, role.name as role_name  FROM "user" u
			LEFT JOIN role ON u.role_id = role.id
		WHERE u.deleted_at IS NULL
	`

	SELECT_USER_BY_ID = `
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, u.deleted_at, role.name as role_name  FROM "user" u
			LEFT JOIN role ON u.role_id = role.id 
		WHERE u.id = $1 AND u.deleted_at IS NULL
	`

	INSERT_USER = `
		INSERT INTO "user" 
			("id", "email", "password", "full_name", "role_id", "created_at", "updated_at", "created_by_id", "updated_by_id")
		VALUES
			($1, $2, $3, $4, $5, NOW(), NOW(), $6, $6)
	`

	UPDATE_USER = `
		UPDATE "user"
			SET ("full_name", "updated_at", "updated_by_id") = ($2, NOW(), $3)
		WHERE id = $1
	`

	HARD_DELETE_USER = `
		DELETE FROM "user"
		WHERE id = $1;
	`

	SOFT_DELETE_USER = `
		UPDATE "user"
			SET ("updated_at", "updated_by_id", "deleted_at", "deleted_by_id") = (NOW(), $2, NOW(), $2)
		WHERE id = $1
	`

	// ADDITIONAL
	SELECT_USER_BY_EMAIL = `
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, role.name as role_name  FROM "user" u
			LEFT JOIN role ON u.role_id = role.id 
		WHERE u.email = $1 AND u.deleted_at IS NULL
	`
)
