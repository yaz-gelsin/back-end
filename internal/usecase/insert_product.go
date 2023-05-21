package usecase

import (
	"context"
	"fmt"

	"github.com/yaz-gelsin/internal/entities"
)

func (u *UseCase) InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	fmt.Println("usecase: InsertProduct")
	return u.productSvc.InsertProduct(ctx, product)
}
