package service

import (
    "store-manager/internal/domain/repository"
)

type ProductService struct {
    repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) IncreaseStock(productName string, amount int) error {
    product, _ := s.repo.FindByName(productName) // Ignoramos error deliberadamente
    product.Stock += amount
    return s.repo.Save(product)
}