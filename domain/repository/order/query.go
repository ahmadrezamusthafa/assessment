package order

type Query string

const (
	QueryInsertOrder Query = `
		INSERT INTO
		order
		(
			id,
			product_id,
			qty,
			created_at
		)
		VALUES
		(
			:id,
			:product_id,
			:qty,
			CURRENT_TIMESTAMP
		)
	`

	QuerySelectOrder Query = `
		SELECT
			id,
			product_id,
			qty,
			created_at,
			updated_at
		FROM order
	`
)

func (c Query) ToString() string {
	return string(c)
}
