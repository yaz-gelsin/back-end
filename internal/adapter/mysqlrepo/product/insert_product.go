package product

import (
	"context"
	"fmt"

	"github.com/yaz-gelsin/internal/entities"
)

func (r *Repository) InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	const queryString = `INSERT INTO
	products 
	(name, description, price, quantity, category_id) 
	VALUES 
	(:name, :description, :price, :quantity, :category_id)`

	_, err := r.conn.NamedQueryContext(ctx, queryString, &product)
	if err != nil {
		return product, fmt.Errorf("error while inserting product: %w", err)
	}
	return product, nil
}
