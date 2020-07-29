package magazinegun

import "time"

type Magazine struct {
	ID         string     `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	BulletQty  int        `db:"bullet_qty" json:"bullet_qty"`
	IsVerified bool       `db:"is_verified" json:"is_verified"`
	Status     int        `db:"status" json:"status"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
}

const (
	StatusAttach = 1
	StatusDetach = 2
)
