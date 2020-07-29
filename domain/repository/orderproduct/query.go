package orderproduct

type Query string

const (
	QueryInsertOrderProduct Query = `
		INSERT INTO
		order_product
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

	QuerySelectOrderProduct Query = `
		SELECT
			id,
			product_id,
			qty,
			created_at,
			updated_at
		FROM order_product
	`
)

func (c Query) ToString() string {
	return string(c)
}
