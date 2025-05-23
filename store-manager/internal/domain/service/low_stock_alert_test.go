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
	err := productSvc.IncreaseStock(productName, increment)

	// Assert (validación)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

}