package product

import (
	"context"
	"fmt"

	"github.com/yaz-gelsin/internal/entities"
)

func (r *Repository) GetByID(ctx context.Context) ([]entities.Product, error) {
	const queryString = `SELECT * FROM
             products 
         WHERE id = ?`

	var products []entities.Product

	err := r.conn.GetContext(ctx, &products, queryString)
	if err != nil {
		return products, fmt.Errorf("error while inserting product: %w", err)
	}
	return products, nil
}
