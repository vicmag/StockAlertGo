package interfaces

import "store-manager/internal/domain/models"

type ProductRepository interface {
	FindByName(name string) (*models.Product, error)
	Save(product *models.Product) error 
}