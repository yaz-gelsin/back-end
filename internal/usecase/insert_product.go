package usecase

import (
	"context"
	"github.com/yaz-gelsin/internal/entities"
)

func (u *UseCase) InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	return u.productSvc.InsertProduct(ctx, product)
}
