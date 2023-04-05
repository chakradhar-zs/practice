package service

import (
	"Day-19/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Product interface {
	GetProduct(ctx *gofr.Context, i int, brand string) (models.Product, error)
	CreateProduct(ctx *gofr.Context, p *models.Product) (interface{}, error)
	UpdateProduct(ctx *gofr.Context, i int, p *models.Product) (interface{}, error)
	DeleteProduct(ctx *gofr.Context, i int) (interface{}, error)
	GetAllProducts(ctx *gofr.Context, brand string) ([]models.Product, error)
	GetProductByNAme(ctx *gofr.Context, name string, brand string) ([]models.Product, error)
}

type Brand interface {
	GetBrand(ctx *gofr.Context, id int) (models.Brand, error)
	CreateBrand(ctx *gofr.Context, brand models.Brand) (interface{}, error)
	UpdateBrand(ctx *gofr.Context, id int, brand models.Brand) (interface{}, error)
}
