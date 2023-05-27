package product

import (
	"context"

	"github.com/yaz-gelsin/internal/entities"
)

func (s *Service) InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	return s.productStore.InsertProduct(ctx, product)
}
