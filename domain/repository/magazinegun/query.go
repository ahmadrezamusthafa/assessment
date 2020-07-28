package magazinegun

type Query string

const (
	QueryInsertMagazine Query = `
		INSERT INTO
		magazine_gun
		(
			id,
			name,
			bullet_qty,
			is_verified,
			status,
			created_at
		)
		VALUES
		(
			:id,
			:name,
			:bullet_qty,
			false,
			:status,
			CURRENT_TIMESTAMP
		)
	`

	QueryUpdateMagazine Query = `
		UPDATE
		magazine_gun
		SET 
			bullet_qty=:bullet_qty,
			is_verified=:is_verified,
			status=:status,
			updated_at=CURRENT_TIMESTAMP
		WHERE
			id=:id;
	`

	QuerySelectMagazine Query = `
		SELECT
			id,
			name,
			bullet_qty,
			is_verified,
			status,
			created_at,
			updated_at
		FROM magazine_gun
	`
)

func (c Query) ToString() string {
	return string(c)
}
