package store

import (
	"api-product/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductStorer interface {
	Get(ctx *gofr.Context, id int, brand string) (interface{}, error)
	Create(ctx *gofr.Context, prod models.Product) (interface{}, error)
	Update(ctx *gofr.Context, id int, prod models.Product) (interface{}, error)
	Del(ctx *gofr.Context, id int) (interface{}, error)
	GetAll(ctx *gofr.Context, brand string) ([]models.Product, error)
}

type BrandStorer interface {
	Get(ctx *gofr.Context, id int) (interface{}, error)
	Create(ctx *gofr.Context, brand models.Brand) (interface{}, error)
	Update(ctx *gofr.Context, id int, brand models.Brand) (interface{}, error)
}
