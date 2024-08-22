package entity

import "time"

type Record struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Price     int       `db:"price"`
	Params    string    `db:"params"`
	Deleted   bool      `db:"is_deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
