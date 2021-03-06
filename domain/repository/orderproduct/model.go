package orderproduct

import "time"

type OrderProduct struct {
	ID        string     `db:"id" json:"id"`
	OrderID   string     `db:"order_id" json:"order_id"`
	ProductID string     `db:"product_id" json:"product_id"`
	Qty       int        `db:"qty" json:"qty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
