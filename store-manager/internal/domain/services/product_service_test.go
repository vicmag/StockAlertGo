package services_test

import (
	"testing"
	"github.com/stretchr/testify/assert"	
	"github.com/stretchr/testify/mock"

	"store-manager/internal/domain/models"
	"store-manager/internal/domain/services"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindByName(name string) (*models.Product, error){
	args := m.Called(name)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Save(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func TestDecrementStock_WhenProductExists_ShouldDecrementStock(t *testing.T){
	//Arrange
	productName := "Camiseta"
	initialStock := 10
	decrement := 5
	finalStock := initialStock - decrement

	product := &models.Product {
		Name: productName,
		Stock: initialStock,
	}

	mockRepo := new(MockProductRepository)

	//Configuración de las expectativas
	mockRepo.On("FindByName", productName).
		Return(product, nil).
		Once()

	mockRepo.On("Save",mock.MatchedBy(func(p *models.Product) bool {
		return p.Stock == finalStock
		})).
		Return(nil).
		Once()

	service := services.NewProductService(mockRepo)

	//Act
	res := service.DecrementStock(productName, decrement)

	//Assert
	assert.NoError(t, res)
	mockRepo.AssertExpectations(t)
}