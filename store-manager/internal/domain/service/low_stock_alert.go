package service

import (
	"store-manager/internal/domain/repository"
	"store-manager/internal/domain/model"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductService{
	return &ProductService{ productRepository: repository}
}

func (s *ProductService) IncrementStock(name string, increment int) error {
	//Implementación de la Fase Verde
	product,_ := s.productRepository.FindByName(name)
	s.updateStock(product, increment)
	s.productRepository.Save(product)
	return nil
}

func (s *ProductService) updateStock(product *model.Product, increment int ){
	product.Stock += increment
}
