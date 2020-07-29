package order

import "time"

type Order struct {
	ID        string     `db:"id" json:"id"`
	ProductID string     `db:"product_id" json:"product_id"`
	Qty       int        `db:"qty" json:"qty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
