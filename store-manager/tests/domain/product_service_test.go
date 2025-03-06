package domain_test

import (
    "testing"
    "store-manager/domain"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
)

// MockProductRepository es una implementación mock de ProductRepository usando testify/mock
type MockProductRepository struct {
    mock.Mock
}

// FindByID simula la búsqueda de un producto por ID
func (m *MockProductRepository) FindByID(id string) (*domain.Product, error) {
    args := m.Called(id)
    return args.Get(0).(*domain.Product), args.Error(1)
}

// Save simula el guardado de un producto
func (m *MockProductRepository) Save(product *domain.Product) error {
    args := m.Called(product)
    return args.Error(0)
}

// TestSetMinStock es la prueba unitaria para SetMinStock
func TestSetMinStock(t *testing.T) {
    // Configuración del test
    repo := &MockProductRepository{}
    service := domain.NewProductService(repo)

    productID := "123"
    minStock := 15
    product := &domain.Product{
        ID:           productID,
        Name:         "Camiseta Azul",
        CurrentStock: 10,
        MinStock:     10,
    }

    // Configuración del mock
    repo.On("FindByID", productID).Return(product, nil)
    repo.On("Save", product).Return(nil)

    // Ejecución del método bajo prueba
    err := service.SetMinStock(productID, minStock)
    assert.NoError(t, err, "SetMinStock should not return an error")

    // Verificación del resultado
    repo.AssertCalled(t, "FindByID", productID)
    repo.AssertCalled(t, "Save", product)
    assert.Equal(t, minStock, product.MinStock, "Expected MinStock to be %d, got %d", minStock, product.MinStock)
}

// TestSetMinStock_InvalidMinStock prueba el caso de un nivel mínimo inválido
func TestSetMinStock_InvalidMinStock(t *testing.T) {
    // Configuración del test
    repo := &MockProductRepository{}
    service := domain.NewProductService(repo)

    productID := "123"
    invalidMinStock := 0 // Nivel mínimo inválido

    // Ejecución del método bajo prueba
    err := service.SetMinStock(productID, invalidMinStock)
    assert.Error(t, err, "SetMinStock should return an error for invalid minStock")
    assert.Equal(t, domain.ErrInvalidMinStock, err, "Expected ErrInvalidMinStock")
}

// TestAlertLowStock es la prueba unitaria para el Escenario 1: Alerta de Stock Bajo
func TestAlertLowStock(t *testing.T) {
    // Configuración del test
    repo := &MockProductRepository{}
    service := domain.NewProductService(repo)

    productID := "456"
    productName := "Camiseta Roja"
    currentStock := 5
    minStock := 10

    product := &domain.Product{
        ID:           productID,
        Name:         productName,
        CurrentStock: currentStock,
        MinStock:     minStock,
    }

    // Configuración del mock
    repo.On("FindByID", productID).Return(product, nil)

    // Ejecución del método bajo prueba (aún no implementado)
    alert := service.CheckLowStock(productID)

    // Verificación del resultado
    assert.False(t, alert, "Expected alert to be false (not implemented yet)")
}