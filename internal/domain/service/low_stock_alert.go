package service

import (
	"store-manager/internal/domain/repository"
)

type LowStockAlertService struct {
	repo repository.ProductRepository
}

func NewLowStockAlertService(repo repository.ProductRepository) *LowStockAlertService {
	return &LowStockAlertService{repo: repo}
}

func (s *LowStockAlertService) CheckLowStock(productID string) (bool, error) {
	// Not implemented yet (Red Phase)
	panic("not implemented")
}
