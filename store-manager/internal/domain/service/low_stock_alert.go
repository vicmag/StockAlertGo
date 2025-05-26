package service

import (
	"store-manager/internal/domain/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService{
	return &ProductService{ productRepository: repo}
}

func (s *ProductService) IncreaseStock(name string, increment int) error {

	return nil
}
