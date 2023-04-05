package product

import (
	"api-product/internal/models"
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) Get(ctx *gofr.Context, id int, brand string) (interface{}, error) {
	var p models.Product

	resp := ctx.DB().QueryRowContext(ctx, "select id,name,description,price,quantity,category,brand_id,status from products where id=?", id)
	fmt.Println(resp)
	err := resp.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.Id, &p.Status)
	if err != nil {
		return models.Product{}, errors.EntityNotFound{}
	}
	if brand == "true" {
		res := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", p.Brand.Id)
		_ = res.Scan(&p.Brand.Name)
	}
	return p, nil
}

func (s *Store) Create(ctx *gofr.Context, prod models.Product) (interface{}, error) {

	resp, err := ctx.DB().ExecContext(ctx, "insert into products values(?,?,?,?,?,?,?,?)", prod.Id, prod.Name, prod.Description, prod.Price, prod.Quantity, prod.Category, prod.Brand.Id, prod.Status)

	if err != nil {
		return int64(0), errors.MissingParam{}
	}
	res, _ := resp.RowsAffected()

	return res, nil
}

func (s *Store) Update(ctx *gofr.Context, id int, prod models.Product) (interface{}, error) {

	resp, err := ctx.DB().ExecContext(ctx, "update products set name=?,description=?,price=?,quantity=?,category=?,brand_id=?,status=? where id =?", prod.Name, prod.Description, prod.Price, prod.Quantity, prod.Category, prod.Brand.Id, prod.Status, id)

	if err != nil {
		return int64(0), errors.EntityNotFound{}
	}

	res, _ := resp.RowsAffected()
	return res, nil
}

func (s *Store) Del(ctx *gofr.Context, id int) (interface{}, error) {
	resp, err := ctx.DB().ExecContext(ctx, "delete from products where id=?", id)

	if err != nil {
		return int64(0), errors.EntityNotFound{Entity: "id"}
	}
	res, _ := resp.RowsAffected()
	return res, nil
}

func (s *Store) GetAll(ctx *gofr.Context, brand string) ([]models.Product, error) {
	res := []models.Product{}
	resp, err := ctx.DB().QueryContext(ctx, "select * from products")

	if err != nil {
		return nil, errors.EntityNotFound{}
	}
	for resp.Next() {
		var p models.Product
		_ = resp.Scan(&p.Id, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.Id, &p.Status)
		if brand == "true" {
			res := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", p.Brand.Id)
			_ = res.Scan(&p.Brand.Name)
		}
		res = append(res, p)

	}
	return res, nil
}
