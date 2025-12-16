package services

import "store-manager/internal/domain/interfaces"


type ProductService struct {
	repo interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository) *ProductService{
	return &ProductService{
		repo: repo, 
	}

}

func (s *ProductService) DecrementStock(productName string, decrement int) error {
	return nil
}