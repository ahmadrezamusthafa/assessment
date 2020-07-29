package product

type Query string

const (
	QueryInsertProduct Query = `
		INSERT INTO
		product
		(
			id,
			code,
			name,
			qty,
			created_at
		)
		VALUES
		(
			:id,
			:code,
			:name,
			:qty,
			CURRENT_TIMESTAMP
		)
	`

	QueryUpdateProduct Query = `
		UPDATE
		product
		SET 
			name=:name,
			qty=:qty,
			updated_at=CURRENT_TIMESTAMP
		WHERE
			id=:id;
	`

	QuerySelectProduct Query = `
		SELECT
			id,
			code,
			name,
			qty,
			created_at,
			updated_at
		FROM product
	`
)

func (c Query) ToString() string {
	return string(c)
}
