package service

import (
	"errors"
	"store-manager/internal/domain/repository"
)

var ErrProductoNoEncontrado = errors.New("producto no encontrado")

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService{
	return &ProductService{ productRepository: repo}
}

func (s *ProductService) IncreaseStock(name string, increment int) error {
	product, _ := s.productRepository.FindByName(name)
	if product == nil {
		return ErrProductoNoEncontrado
	}
	product.Stock += increment
	s.productRepository.Save(product)
	return nil
}
