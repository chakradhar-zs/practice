package brand

import (
	"reflect"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
	"Day-19/internal/store"
)

type Service struct {
	store store.BrandStorer
}

func New(storer store.BrandStorer) *Service {
	return &Service{store: storer}
}

func (svc *Service) GetBrand(ctx *gofr.Context, id int) (models.Brand, error) {
	res, err := svc.store.Get(ctx, id)

	if err != nil {
		return models.Brand{}, err
	}

	return res, nil
}

func (svc *Service) CreateBrand(ctx *gofr.Context, brand models.Brand) (interface{}, error) {
	vals := reflect.ValueOf(brand)

	for i := 0; i < vals.NumField(); i++ {
		if vals.Field(i).Interface() == "" || vals.Field(i).Interface() == 0 {
			return 0, errors.MissingParam{Param: []string{"body"}}
		}
	}

	res, _ := svc.store.Create(ctx, brand)

	return res, nil
}

func (svc *Service) UpdateBrand(ctx *gofr.Context, id int, brand models.Brand) (interface{}, error) {
	vals := reflect.ValueOf(brand)

	for i := 0; i < vals.NumField(); i++ {
		if vals.Field(i).Interface() == "" || vals.Field(i).Interface() == 0 {
			return 0, errors.MissingParam{Param: []string{"body"}}
		}
	}

	res, err := svc.store.Update(ctx, id, brand)

	if err != nil {
		return 0, err
	}

	return res, nil
}
