package services

import "store-manager/internal/domain/interfaces"

type ProductService struct {
	repo interfaces.ProductRepository
}

func NewProductService(db interfaces.ProductRepository) *ProductService {
	return &ProductService{
		repo: db,
	}
}

func (s *ProductService) IncrementStock(name string, increment int) error {
	product, _ := s.repo.FindByName(name)

	product.Stock += increment

	_ = s.repo.Save(product)

	return nil
}
