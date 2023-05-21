package product

import (
	"context"
	"fmt"

	"github.com/yaz-gelsin/internal/entities"
)

func (r *Repository) InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	fmt.Println("repo: InsertProduct")
	const queryString = `INSERT INTO
	products 
	(name, description, price, quantity, category_id, created_at, updated_at) 
	VALUES 
	(:name, :description, :price, :quantity, :category_id, :created_at, :updated_at)`

	_, err := r.conn.NamedExecContext(ctx, queryString, product)
	if err != nil {
		return product, fmt.Errorf("error while inserting product: %w", err)
	}
	return product, nil
}
