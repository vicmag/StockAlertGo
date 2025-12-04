package repository

import "store-manager/internal/domain/model"

type ProductRepository interface {
	FindByName(name string) (*model.Product, error)
	Save(product *model.Product) error
}