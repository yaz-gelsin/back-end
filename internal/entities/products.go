package entities

import (
	"gopkg.in/guregu/null.v4"
)

type Product struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	Quantity    int64     `db:"quantity" json:"quantity"`
	CategoryID  int64     `db:"category_id" json:"category_id"`
	CreatedAt   null.Time `db:"created_at" json:"created_at"`
	UpdatedAt   null.Time `db:"updated_at" json:"updated_at"`
}
