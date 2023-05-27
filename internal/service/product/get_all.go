package product

import (
	"context"
	"github.com/yaz-gelsin/internal/entities"
)

func (s *Service) GetAll(ctx context.Context) ([]entities.Product, error) {
	return s.productStore.GetAll(ctx)
}
