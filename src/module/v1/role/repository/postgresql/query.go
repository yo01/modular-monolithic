package postgresql

const (
	SELECT_ROLE = `
		SELECT * FROM "role" r
		WHERE r.deleted_at IS NULL
	`

	SELECT_ROLE_BY_ID = `
		SELECT * FROM "role" r WHERE r.id = $1
		WHERE r.deleted_at IS NULL
	`

	INSERT_ROLE = `
		INSERT INTO "role" 
			("id", "name", "created_at", "updated_at")
		VALUES
			($1, $2, NOW(), NOW())
	`

	UPDATE_ROLE = `
		UPDATE "role"
			SET ("name", "updated_at", "updated_by_id") = ($2, NOW(), $3)
		WHERE id = $1
	`

	HARD_DELETE_ROLE = `
		DELETE FROM "role"
		WHERE id = $1;
	`

	SOFT_DELETE_ROLE = `
		UPDATE "role"
			SET ("updated_at", "updated_by_id", "deleted_at", "deleted_by_id") = (NOW(), $2, NOW(), $2)
		WHERE id = $1
	`
)
