package repository

import "store-manager/internal/domain/model"

type ProductRepository interface {
	FindByID(id string) (*model.Product, error)
}
