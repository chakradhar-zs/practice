package product

import (
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
)

type Store struct {
}

func New() *Store {
	return &Store{}
}

const val = "true"

func (s *Store) Get(ctx *gofr.Context, id int, brand string) (models.Product, error) {
	var p models.Product

	resp := ctx.DB().QueryRowContext(ctx, "select id,name,description,price,quantity,category,brand_id,status from products where id=?", id)
	fmt.Println(resp)
	err := resp.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.ID, &p.Status)

	if err != nil {
		return models.Product{}, errors.EntityNotFound{}
	}

	if brand == val {
		res := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", p.Brand.ID)
		_ = res.Scan(&p.Brand.Name)
	}

	return p, nil
}

func (s *Store) GetByName(ctx *gofr.Context, name, brand string) ([]models.Product, error) {
	res := []models.Product{}

	resp, _ := ctx.DB().QueryContext(ctx,
		"select id,name,description,price,quantity,category,brand_id,status from products where name=?",
		name)
	if resp == nil {
		return []models.Product{{}}, errors.EntityNotFound{}
	}

	for resp.Next() {
		var p models.Product
		_ = resp.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.ID, &p.Status)

		if brand == val {
			res := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", p.Brand.ID)
			_ = res.Scan(&p.Brand.Name)
		}

		res = append(res, p)
	}

	return res, nil
}

func (s *Store) Create(ctx *gofr.Context, prod *models.Product) (interface{}, error) {
	resp, err := ctx.DB().ExecContext(ctx, "insert into products values(?,?,?,?,?,?,?,?)",
		prod.ID, prod.Name, prod.Description, prod.Price, prod.Quantity, prod.Category, prod.Brand.ID, prod.Status)

	if err != nil {
		return int64(0), errors.MissingParam{}
	}

	res, _ := resp.RowsAffected()

	return res, nil
}

func (s *Store) Update(ctx *gofr.Context, id int, prod *models.Product) (interface{}, error) {
	resp, err := ctx.DB().ExecContext(ctx,
		"update products set name=?,description=?,price=?,quantity=?,category=?,brand_id=?,status=? where id =?",
		prod.Name, prod.Description, prod.Price, prod.Quantity, prod.Category, prod.Brand.ID, prod.Status, id)

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
		_ = resp.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity, &p.Category, &p.Brand.ID, &p.Status)

		if brand == val {
			res := ctx.DB().QueryRowContext(ctx, "select name from brands where id=?", p.Brand.ID)
			_ = res.Scan(&p.Brand.Name)
		}

		res = append(res, p)
	}

	return res, nil
}
