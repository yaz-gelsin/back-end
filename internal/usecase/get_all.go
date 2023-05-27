package usecase

import (
	"context"
	"github.com/yaz-gelsin/internal/entities"
)

func (u *UseCase) GetAll(ctx context.Context) ([]entities.Product, error) {
	return u.productSvc.GetAll(ctx)
}
