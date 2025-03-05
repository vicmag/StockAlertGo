package domain_test

import (
    "testing"
    "store-manager/domain"
    "store-manager/infrastructure/mocks"
)

// TestSetMinStock es la prueba unitaria para SetMinStock
func TestSetMinStock(t *testing.T) {
    // Configuración del test
    repo := &mocks.MockProductRepository{}
    service := domain.NewProductService(repo)

    productID := "123"
    minStock := 15

    // Ejecución del método bajo prueba
    err := service.SetMinStock(productID, minStock)
    if err != nil {
        t.Fatalf("SetMinStock failed: %v", err)
    }

    // Verificación del resultado
    product, err := repo.FindByID(productID)
    if err != nil {
        t.Fatalf("FindByID failed: %v", err)
    }
    if product.MinStock != minStock {
        t.Errorf("Expected MinStock to be %d, got %d", minStock, product.MinStock)
    }
}
    