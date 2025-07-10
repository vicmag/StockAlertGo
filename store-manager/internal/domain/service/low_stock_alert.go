package service

import (
	"store-manager/internal/domain/repository"
	"errors"
)

const errProductNotFound = "producto no encontrado"
const errIncrementMustBePositive = "incremento debe ser positivo"


type ProductService struct{
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService{
	return &ProductService{productRepository: repo}

}

func (s *ProductService) IncrementStock(name string, increment int) error{		
	if increment <= 0 {
		return errors.New(errIncrementMustBePositive)
	}	

	product, _ := s.productRepository.FindByName(name)		
	if product == nil {
		return errors.New(errProductNotFound)
	}

	product.Stock += increment	
	s.productRepository.Save(product)

	return nil
}