package store

import "mock-productservice/internal/models"

type ProductStorer interface {
	Create(m *models.Product) (*models.Product, error)
}
