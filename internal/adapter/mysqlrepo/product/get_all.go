package product

import (
	"context"
	"fmt"

	"github.com/yaz-gelsin/internal/entities"
)

func (r *Repository) GetAll(ctx context.Context) ([]entities.Product, error) {
	const queryString = `SELECT * FROM products`

	var products []entities.Product

	err := r.conn.SelectContext(ctx, &products, queryString)
	if err != nil {
		return products, fmt.Errorf("error while inserting product: %w", err)
	}
	return products, nil
}
