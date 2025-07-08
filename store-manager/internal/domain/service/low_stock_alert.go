package service

import "store-manager/internal/domain/repository"

type ProductService struct{
	productRepository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService{
	return &ProductService{productRepository: repo}

}

func (s *ProductService) IncrementStock(nameame string, increment int) error{
	//Impletanción vacia. Fase Roja
	return nil
}