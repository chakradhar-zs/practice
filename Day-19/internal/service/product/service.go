package product

import (
	"reflect"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Day-19/internal/models"
	"Day-19/internal/store"
)

type Service struct {
	store store.ProductStorer
}

func New(storer store.ProductStorer) *Service {
	return &Service{store: storer}
}

func (s *Service) GetProduct(ctx *gofr.Context, i int, brand string) (models.Product, error) {
	res, err := s.store.Get(ctx, i, brand)

	if err != nil {
		return models.Product{}, err
	}

	return res, nil
}

func (s *Service) GetProductByNAme(ctx *gofr.Context, name, brand string) ([]models.Product, error) {
	res, err := s.store.GetByName(ctx, name, brand)

	if err != nil {
		return []models.Product{}, err
	}

	return res, nil
}

func (s *Service) CreateProduct(ctx *gofr.Context, p *models.Product) (interface{}, error) {
	x := *p
	values := reflect.ValueOf(x)

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == "" || values.Field(i).Interface() == 0 {
			return 0, errors.MissingParam{Param: []string{"body"}}
		}
	}

	res, _ := s.store.Create(ctx, p)

	return res, nil
}

func (s *Service) UpdateProduct(ctx *gofr.Context, id int, p *models.Product) (interface{}, error) {
	x := *p
	values := reflect.ValueOf(x)

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == "" || values.Field(i).Interface() == 0 {
			return 0, errors.MissingParam{Param: []string{"body"}}
		}
	}

	res, err := s.store.Update(ctx, id, p)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *Service) DeleteProduct(ctx *gofr.Context, i int) (interface{}, error) {
	res, err := s.store.Del(ctx, i)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *Service) GetAllProducts(ctx *gofr.Context, brand string) ([]models.Product, error) {
	res, err := s.store.GetAll(ctx, brand)

	if err != nil {
		return []models.Product{}, err
	}

	return res, nil
}
