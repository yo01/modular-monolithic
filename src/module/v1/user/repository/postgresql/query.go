package postgresql

const (
	SELECT_USER = `
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, u.created_at, u.created_at, u.deleted_at, role.name as role_name, role.id as role_id, p.id as permission_id, p.role_id as permission_role_id, p.list_api as list_api FROM "user" u
			LEFT JOIN role ON u.role_id = role.id
			LEFT JOIN permission p ON p.role_id = role.id
		WHERE u.deleted_at IS NULL
	`

	SELECT_USER_BY_ID = `
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, u.created_at, u.created_at, u.deleted_at, role.name as role_name, role.id as role_id, p.id as permission_id, p.role_id as permission_role_id, p.list_api as list_api FROM "user" u
			LEFT JOIN role ON u.role_id = role.id
			LEFT JOIN permission p ON p.role_id = role.id
		WHERE u.deleted_at IS NULL AND u.id = $1
	`

	INSERT_USER = `
		INSERT INTO "user" 
			("id", "email", "password", "full_name", "role_id", "created_at", "updated_at")
		VALUES
			($1, $2, $3, $4, $5, NOW(), NOW())
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
		SELECT u.id , u.email , u.password, u.full_name, u.role_id, u.created_at, u.created_at, u.deleted_at, role.name as role_name, role.id as role_id, p.id as permission_id, p.role_id as permission_role_id, p.list_api as list_api FROM "user" u
			LEFT JOIN role ON u.role_id = role.id
			LEFT JOIN permission p ON p.role_id = role.id
		WHERE u.deleted_at IS NULL AND u.email = $1
	`
)
