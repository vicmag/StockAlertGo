package service

import "store-manager/internal/domain/repository" 

type ProductService struct {
	productoDB repository.ProductoDB
}

func NewProductService(db repository.ProductoDB) *ProductService{
	return &ProductService{productoDB: db}
}

func (s *ProductService) IncrementaStock(nombre string, incremento int) error {
	//Fase Verde
	producto,_ := s.productoDB.FindByName(nombre)
	producto.Stock += incremento
	s.productoDB.Save(producto)

	return nil
}