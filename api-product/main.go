package main

import (
	"api-product/internal/http/brand"
	"api-product/internal/http/product"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	productstore "api-product/internal/store/product"

	productservice "api-product/internal/service/product"

	brandservice "api-product/internal/service/brand"
	brandstore "api-product/internal/store/brand"
)

func main() {
	app := gofr.New()

	app.Server.ValidateHeaders = false

	productStore := productstore.New()
	productSvc := productservice.New(productStore)
	prodHTTP := product.New(productSvc)

	brandStore := brandstore.New()
	brandSvc := brandservice.New(brandStore)
	brandHTTP := brand.New(brandSvc)

	app.REST("product", prodHTTP)
	app.REST("brand", brandHTTP)
	app.Start()
}
