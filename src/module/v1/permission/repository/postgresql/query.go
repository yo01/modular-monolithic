package postgresql

const (
	SELECT_PERMISSION = `
		SELECT p.id, p.list_api, p.created_at, p.role_id as role_id, r.name as role_name FROM "permission" p
			LEFT JOIN role r ON r.id = role_id
		WHERE p.deleted_at IS NULL
	`

	SELECT_PERMISSION_BY_ID = `
		SELECT p.id, p.list_api, p.created_at, p.role_id as role_id, r.name as role_name FROM "permission" p 
			LEFT JOIN role r ON r.id = role_id
		WHERE p.deleted_at IS NULL AND p.id = $1
	`

	INSERT_PERMISSION = `
		INSERT INTO "permission" 
			("id", "role_id", "list_api", "created_at", "updated_at")
		VALUES
			($1, $2, $3, NOW(), NOW())
	`

	UPDATE_PERMISSION = `
		UPDATE "permission"
			SET ("role_id", "list_api", "updated_at", "updated_by_id") = ($2, $3, NOW(), $4)
		WHERE id = $1
	`

	HARD_DELETE_PERMISSION = `
		DELETE FROM "permission"
		WHERE id = $1;
	`

	SOFT_DELETE_PERMISSION = `
		UPDATE "permission"
			SET ("updated_at", "updated_by_id", "deleted_at", "deleted_by_id") = (NOW(), $2, NOW(), $2)
		WHERE id = $1
	`

	SELECT_PERMISSION_BY_ROLE_ID = `
		SELECT p.id, p.list_api, p.created_at, p.role_id as role_id, r.name as role_name FROM "permission" p 
			LEFT JOIN role r ON r.id = role_id
		WHERE p.deleted_at IS NULL AND p.role_id = $1
	`
)
