package postgresql

const (
	SELECT_PERMISSION = `
		SELECT * FROM "permission" p
		WHERE p.deleted_at IS NULL
	`

	SELECT_PERMISSION_BY_ID = `
		SELECT * FROM "permission" p WHERE p.id = $1
		WHERE p.deleted_at IS NULL
	`

	INSERT_PERMISSION = `
		INSERT INTO "permission" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_PERMISSION = `
		UPDATE "permission"
			SET ("name", "updated_at", "updated_by_id") = ($2, NOW(), $3)
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
)
