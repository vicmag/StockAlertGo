package service

import "store-manager/internal/domain/repository" 

type ProductService struct {
	productoDB repository.ProductoDB
}

func NewProductService(db repository.ProductoDB) *ProductService{
	return &ProductService{productoDB: db}
}

func (s *ProductService) IncrementaStock(nombre string, incremento int) error {
	//Implementación vacia. Fase Roja
	return nil
}