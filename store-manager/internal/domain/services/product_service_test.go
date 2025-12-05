package services_test

import (
    "testing"
   
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
   
    "store-manager/internal/domain/models"
    "store-manager/internal/domain/services"
)

// MockProductRepository - MÍNIMO: solo lo que necesita el test
type MockProductRepository struct {
    mock.Mock
}

func (m *MockProductRepository) FindByName(name string) (*models.Product, error) {
    args := m.Called(name)
    return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Save(product *models.Product) error {
    args := m.Called(product)
    return args.Error(0)
}

// NUEVO: Mock para servicio de notificación (implementación directa)
type MockNotifierService struct {
    mock.Mock
}

func (m *MockNotifierService) SendLowStockAlert(productName string, currentStock, minStockLevel int) error {
    args := m.Called(productName, currentStock, minStockLevel)
    return args.Error(0)
}

// Test para ESCENARIO 4: Decremento Simple
func TestDecrementStock_WhenProductExists_ShouldDecrementStock(t *testing.T) {
    // Arrange - Configuración exacta del escenario
    productName := "Camiseta"
    initialStock := 10
    decrementAmount := 5
    expectedStock := initialStock - decrementAmount // 5
   
    // Producto mínimo - solo campos necesarios
    mockProduct := &models.Product{
        Name:  productName,
        Stock: initialStock,
    }
   
    mockRepo := new(MockProductRepository)
   
    // Configuración mínima de expectativas
    mockRepo.On("FindByName", productName).
        Return(mockProduct, nil).
        Once()
   
    mockRepo.On("Save", mock.MatchedBy(func(p *models.Product) bool {
        return p.Stock == expectedStock // Stock debe ser 5
    })).
        Return(nil).
        Once()
   
    service := services.NewProductServiceWithoutNotifier(mockRepo)
   
    // Act
    err := service.DecrementStock(productName, decrementAmount)
   
    // Assert
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

func TestDecrementStock_WhenStockFallsBelowMinimum_ShouldSendAlert(t *testing.T) {
    // Given - Configuración del escenario
    productName := "Camiseta Azul"
    initialStock := 15
    minStockLevel := 10
    decrementAmount := 10
    expectedStock := initialStock - decrementAmount // 5 unidades
   
    // Producto con nivel mínimo
    mockProduct := &models.Product{
        Name:          productName,
        Stock:         initialStock,
        MinStockLevel: minStockLevel, // ← ERROR de compilación (campo no existe)
    }
   
    mockRepo := new(MockProductRepository)
    mockNotifier := new(MockNotifierService) // ← NUEVA DEPENDENCIA DIRECTA
   
    // Configurar expectativas del repositorio
    mockRepo.On("FindByName", productName).
        Return(mockProduct, nil).
        Once()
   
    mockRepo.On("Save", mock.MatchedBy(func(p *models.Product) bool {
        return p.Stock == expectedStock
    })).
        Return(nil).
        Once()
   
    // Configurar expectativas del notificador
    // Debe notificar porque stock (5) < minStockLevel (10)
    mockNotifier.On("SendLowStockAlert", productName, expectedStock, minStockLevel).
        Return(nil).
        Once()
   
    // Crear servicio con notificador (ERROR: constructor no acepta notificador)
    service := services.NewProductService(mockRepo, mockNotifier)
   
    // When
    err := service.DecrementStock(productName, decrementAmount)
   
    // Then
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
    mockNotifier.AssertExpectations(t) // ← NUEVA verificación
}