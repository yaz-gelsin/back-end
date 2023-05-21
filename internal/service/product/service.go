package product

import (
	"context"

	"github.com/yaz-gelsin/internal/entities"
)

type productStore interface {
	InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error)
}

type Service struct {
	productStore productStore
}

func NewProductRepo(productStore productStore) *Service {
	return &Service{
		productStore: productStore,
	}
}
