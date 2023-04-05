package brand

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
)

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) Get(ctx *gofr.Context, id int) (models.Brand, error) {
	var b models.Brand

	resp := ctx.DB().QueryRowContext(ctx, "select id,name from brands where id=?", id)

	err := resp.Scan(&b.ID, &b.Name)

	if err != nil {
		return models.Brand{}, errors.EntityNotFound{}
	}

	return b, nil
}
func (s *Store) Create(ctx *gofr.Context, brand models.Brand) (interface{}, error) {
	resp, err := ctx.DB().ExecContext(ctx, "insert into brands values(?,?)", brand.ID, brand.Name)

	if err != nil {
		return int64(0), errors.MissingParam{}
	}

	res, _ := resp.RowsAffected()

	return res, nil
}
func (s *Store) Update(ctx *gofr.Context, id int, brand models.Brand) (interface{}, error) {
	resp, err := ctx.DB().ExecContext(ctx, "update brands set name=? where id=?", brand.Name, id)

	if err != nil {
		return int64(0), errors.EntityNotFound{}
	}

	res, _ := resp.RowsAffected()

	return res, nil
}
