package usecase

import (
	"context"

	"github.com/yaz-gelsin/internal/entities"
)

type productService interface {
	InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error)
	GetAll(ctx context.Context) ([]entities.Product, error)
}

type UseCase struct {
	productSvc productService
}

func NewUseCase(productSvc productService) *UseCase {
	return &UseCase{
		productSvc: productSvc,
	}
}
