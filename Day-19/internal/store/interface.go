package store

import (
	"Day-19/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type ProductStorer interface {
	Get(ctx *gofr.Context, id int, brand string) (models.Product, error)
	Create(ctx *gofr.Context, prod *models.Product) (interface{}, error)
	Update(ctx *gofr.Context, id int, prod *models.Product) (interface{}, error)
	Del(ctx *gofr.Context, id int) (interface{}, error)
	GetAll(ctx *gofr.Context, brand string) ([]models.Product, error)
	GetByName(ctx *gofr.Context, name string, brand string) ([]models.Product, error)
}

type BrandStorer interface {
	Get(ctx *gofr.Context, id int) (models.Brand, error)
	Create(ctx *gofr.Context, brand models.Brand) (interface{}, error)
	Update(ctx *gofr.Context, id int, brand models.Brand) (interface{}, error)
}
