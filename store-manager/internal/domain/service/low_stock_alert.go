package service

import (
	"errors"
	"store-manager/internal/domain/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService{
	return &ProductService{ productRepository: repo}
}

func (s *ProductService) IncreaseStock(name string, increment int) error {
	product, _ := s.productRepository.FindByName(name)
	if product == nil {
		return errors.New("producto no encontrado")
	}
	product.Stock += increment
	s.productRepository.Save(product)
	return nil
}
