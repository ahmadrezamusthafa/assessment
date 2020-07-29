package product

import "time"

type Product struct {
	ID        string     `db:"id" json:"id"`
	Code      string     `db:"code" json:"code"`
	Name      string     `db:"name" json:"name"`
	Qty       int        `db:"qty" json:"qty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
