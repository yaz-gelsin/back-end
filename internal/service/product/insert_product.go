package product

import (
	"context"
	"gopkg.in/guregu/null.v4"
	"time"

	"github.com/yaz-gelsin/internal/entities"
)

func (s *Service) InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	if product.CreatedAt.NullTime.Time.IsZero() {
		product.CreatedAt = null.NewTime(time.Now().UTC(), true)
	}
	product.UpdatedAt = null.NewTime(time.Now().UTC(), true)
	return s.productStore.InsertProduct(ctx, product)
}
