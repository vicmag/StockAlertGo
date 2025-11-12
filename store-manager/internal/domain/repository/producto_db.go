package repository

import "store-manager/internal/domain/model"

type ProductoDB interface {
	FindByName(nombre string) (*model.Producto, error)
	Save(producto *model.Producto) error
}