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
        // 1. Buscar producto por nombre - sin manejo de error
    product, _ := s.repo.FindByName(productName)
   
    // 2. Decrementar el stock
    product.Stock -= amount
   
    // 3. Guardar el producto - sin manejo de error
    _ = s.repo.Save(product)
   
    return nil  // ← Siempre éxito, como el test espera
}