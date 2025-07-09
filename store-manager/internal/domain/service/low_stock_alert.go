package service

import (
	"store-manager/internal/domain/repository"
	"errors"
)

type ProductService struct{
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService{
	return &ProductService{productRepository: repo}

}

func (s *ProductService) IncrementStock(name string, increment int) error{
	//Fase Verde
	product, _ := s.productRepository.FindByName(name)	
	if product == nil {
		return errors.New("producto no encontrado")
	}
	product.Stock += increment	
	s.productRepository.Save(product)

	return nil
}