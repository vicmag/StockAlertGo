package mocks

import (
    "store-manager/domain"
)

// MockProductRepository es una implementación mock de ProductRepository para pruebas
type MockProductRepository struct {
    products map[string]*domain.Product
}

// NewMockProductRepository crea una nueva instancia de MockProductRepository
func NewMockProductRepository() *MockProductRepository {
    return &MockProductRepository{
        products: make(map[string]*domain.Product),
    }
}

// FindByID simula la búsqueda de un producto por ID
func (m *MockProductRepository) FindByID(id string) (*domain.Product, error) {
    if product, exists := m.products[id]; exists {
        return product, nil
    }
    return nil, nil // Simula que no se encontró el producto
}

// Save simula el guardado de un producto
func (m *MockProductRepository) Save(product *domain.Product) error {
    m.products[product.ID] = product
    return nil
}
    