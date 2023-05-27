package product

import "github.com/jmoiron/sqlx"

type Repository struct {
	conn *sqlx.DB
}

func NewProductRepo(conn *sqlx.DB) *Repository {
	return &Repository{
		conn: conn,
	}
}
