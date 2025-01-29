package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"store-manager/internal/domain/model"
	"store-manager/internal/domain/repository"
)

type mockProductRepository struct {
	repository.ProductRepository
	product *model.Product
	err     error
}

func (m *mockProductRepository) FindByID(id string) (*model.Product, error) {
	return m.product, m.err
}

func TestLowStockAlertService_CheckLowStock(t *testing.T) {
	tests := []struct {
		name           string
		product        *model.Product
		expectedResult bool
	}{
		{
			name:           "Stock is low",
			product:        &model.Product{ID: "1", Name: "Product A", CurrentStock: 5, LowStockThreshold: 10},
			expectedResult: true,
		},
		{
			name:           "Stock is not low",
			product:        &model.Product{ID: "1", Name: "Product A", CurrentStock: 15, LowStockThreshold: 10},
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockProductRepository{product: tt.product}
			service := NewLowStockAlertService(repo)

			result, err := service.CheckLowStock("1")
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}