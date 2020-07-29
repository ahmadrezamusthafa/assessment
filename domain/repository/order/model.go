package order

import "time"

type Order struct {
	ID         string     `db:"id" json:"id"`
	IsVerified bool       `db:"is_verified" json:"is_verified"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
}
