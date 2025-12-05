package services

import "store-manager/internal/domain/ports"

// ProductService - Estructura mínima
type ProductService struct {
    repo ports.ProductRepository
}

// NewProductService - Constructor mínimo
func NewProductService(repo ports.ProductRepository) *ProductService {
    return &ProductService{
        repo: repo,
    }
}

// DecrementStock - Implementación MÍNIMA para fase roja
// Panic porque aún no está implementado
func (s *ProductService) DecrementStock(productName string, amount int) error {
    panic("DecrementStock no implementado - Fase Roja")
}