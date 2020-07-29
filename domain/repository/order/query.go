package order

type Query string

const (
	QueryInsertOrder Query = `
		INSERT INTO
		order
		(
			id,
			is_verified,
			created_at
		)
		VALUES
		(
			:id,
			false,
			CURRENT_TIMESTAMP
		)
	`

	QueryUpdateOrder Query = `
		UPDATE
		order
		SET 
			is_verified=:is_verified,
			updated_at=CURRENT_TIMESTAMP
		WHERE
			id=:id;
	`

	QuerySelectOrder Query = `
		SELECT
			id,
			is_verified,
			created_at,
			updated_at
		FROM order
	`
)

func (c Query) ToString() string {
	return string(c)
}
