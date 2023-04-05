package main

import (
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
	Quantity    int    `json:"q"`
	Category    string `json:"category"`
	Brand_id    int    `json:"brandid"`
	Status      string `json:"status"`
}

func RetrieveProduct(ctx *gofr.Context, id int) (Product, error) {

	var p1 Product

	row := ctx.DB().QueryRow("select * from products where id=?", id)

	if row.Err() != nil {
		return Product{}, errors.InvalidParam{}
	}

	err := row.Scan(&p1.Id, &p1.Name, &p1.Description, &p1.Price, &p1.Quantity, &p1.Category, &p1.Brand_id, &p1.Status)

	if err != nil {
		return Product{}, errors.EntityNotFound{}
	}

	return p1, nil
}

func Create(ctx *gofr.Context, p Product) (int64, error) {

	result, err := ctx.DB().Exec("insert into products values(?,?,?,?,?,?,?,?)", p.Id, p.Name, p.Description, p.Price, p.Quantity, p.Category, p.Brand_id, p.Status)

	if err != nil {
		return 0, errors.MissingParam{}
	}
	row, _ := result.RowsAffected()

	return row, nil
}

func getByID(ctx *gofr.Context) (interface{}, error) {

	str := ctx.PathParam("id")

	if str == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		return Product{}, errors.InvalidParam{}
	}

	out, err := RetrieveProduct(ctx, id)
	return out, err
}

func CreateProduct(ctx *gofr.Context) (interface{}, error) {
	var p Product
	err := ctx.Bind(&p)
	if err != nil {
		return int64(0), errors.MissingParam{}
	}

	resp, err := Create(ctx, p)
	if err != nil {
		return int64(0), err
	}
	return resp, nil
}

func main() {
	app := gofr.New()

	app.Server.ValidateHeaders = false

	app.GET("/product/{id}", getByID)

	app.POST("/product", CreateProduct)

	app.Start()

}
