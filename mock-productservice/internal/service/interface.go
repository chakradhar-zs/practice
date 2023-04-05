package service

type BrandService interface {
	Get(name string) (int, error)
	Create(name string) (int, error)
}
