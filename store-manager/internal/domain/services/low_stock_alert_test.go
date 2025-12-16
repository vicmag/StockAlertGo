package services_test

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
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

func TestDecrementStock_DeberiaGuardarElProducto_AlIncrementar(t *testing.T){
	//Arrange (Configuración)
	productName := "Camiseta"
	initialStock := 10
	decrement := 5
	expectedStock := initialStock - decrement

	product := &models.Product{
		Name: productName,
		Stock: initialStock,
	}

	mockRepo := new(MockProductRepository)
	mockRepo.On("FindByName", productName).
		Return(product, nil).
		Once()

	mockRepo.On("Save", mock.MatchedBy(func(p *models.Product) bool{
			return p.Stock == expectedStock
		})).
		Return(nil).
		Once()

	service := services.NewProductService(mockRepo)

	//Act (Ejecución)
	res := service.DecrementStock(productName, decrement)

	//Assert (Validación)
	assert.NoError(t, res)
	mockRepo.AssertExpectations(t)
}