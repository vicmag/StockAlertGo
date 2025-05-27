package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"store-manager/internal/domain/model"
	"store-manager/internal/domain/service"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindByName(name string) (*model.Product, error){
	args := m.Called(name)
	if args.Get(0) ==nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductRepository) Save(product *model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}


func TestIncrementStock(t *testing.T){
	// Arrange (configuración)
	mockRepo := new(MockProductRepository)
	productSvc := service.NewProductService(mockRepo)
	productName := "Camiseta"
	initialStock := 10
	increment := 5

	product := &model.Product{
		Name:   productName,
		Stock:  initialStock,
	}

	//Configuro las expectativas del mock
	mockRepo.On("FindByName", productName).
		Return(product, nil).
		Once()

	mockRepo.On("Save", mock.AnythingOfType("*model.Product")).
		Run(func(args mock.Arguments) {
			p := args.Get(0).(*model.Product)
			assert.Equal(t, initialStock+increment, p.Stock)

		}).
		Return(nil).
		Once()

	 
	// Act (ejecución)
	productSvc.IncreaseStock(productName, increment)

	// Assert (validación)
	mockRepo.AssertExpectations(t)

}

func TestIncrementStock_ProductoNoEncontrado(t *testing.T){
	// Arrange (configuración)
	mockRepo := new(MockProductRepository)
	productSvc := service.NewProductService(mockRepo)
	productName := "ProductoInexistente"

	// Configuración del mock
	mockRepo.On("FindByName", productName).
		Return(nil, nil).
		Once()

	// Act (ejecución)
	err := productSvc.IncreaseStock(productName, 5)

	// Assert (validación)
	assert.Error(t, err) // Verifica que se retorne un error
	assert.Equal(t, "producto no encontrado", err.Error())
	mockRepo.AssertExpectations(t) // Verifica que se llamaron las expectativas del mock

}
