package service

import ("store-manager/internal/domain/repository")

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService{
	return &ProductService{productRepository: repo}
}

func (s *ProductService) IncrementStock(name string, increment int) error {
	product, _ := s.productRepository.FindByName(name)
	product.Stock += increment
	s.productRepository.Save(product)

	return nil
}