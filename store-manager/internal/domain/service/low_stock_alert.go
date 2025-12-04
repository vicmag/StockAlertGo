package service

import "store-manager/internal/domain/repository"

type ProductService struct{
	productDB repository.ProductRepository	
}

func NewProductService(db repository.ProductRepository) *ProductService{
	return &ProductService{ productDB: db}
}


func (s* ProductService) IncrementStock(name string, increment int) error {
	return nil
}