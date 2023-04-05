package product

import (
	"errors"
	"mock-productservice/internal/models"
	"mock-productservice/internal/service"
	"mock-productservice/internal/store"
)

type Service struct {
	store    store.ProductStorer
	brandSvc service.BrandService
}

func New(storer store.ProductStorer, brandSvc service.BrandService) *Service {
	return &Service{
		store:    storer,
		brandSvc: brandSvc,
	}
}

func (s *Service) Create(req *service.ProductCreate) (*models.Product, error) {
	if len(req.Name) < 3 {
		return nil, errors.New("invalid product name")
	}

	brandId, err := s.brandSvc.Get(req.BrandName)
	if err != nil {
		return nil, err
	}
	if brandId == 0 {
		brandId, err = s.brandSvc.Create(req.BrandName)
		if err != nil {
			return nil, err
		}
	}
	product, err := s.store.Create(&models.Product{
		Name:    req.Name,
		BrandId: brandId,
	})
	return product, err
}
