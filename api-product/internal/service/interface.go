package service

import (
	"api-product/internal/models"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Product interface {
	GetProduct(ctx *gofr.Context, i int, brand string) (interface{}, error)
	CreateProduct(ctx *gofr.Context, p models.Product) (interface{}, error)
	UpdateProduct(ctx *gofr.Context, i int, p models.Product) (interface{}, error)
	DeleteProduct(ctx *gofr.Context, i int) (interface{}, error)
	GetAllProducts(ctx *gofr.Context, brand string) (interface{}, error)
}

type Brand interface {
	GetBrand(ctx *gofr.Context, id int) (interface{}, error)
	CreateBrand(ctx *gofr.Context, brand models.Brand) (interface{}, error)
	UpdateBrand(ctx *gofr.Context, id int, brand models.Brand) (interface{}, error)
	//DeleteBrand(ctx *gofr.Context, id int) (interface{}, error)
}
