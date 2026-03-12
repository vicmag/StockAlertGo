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

func (s *ProductService) IncrementStock(name string, increment int) error{
	return nil
}