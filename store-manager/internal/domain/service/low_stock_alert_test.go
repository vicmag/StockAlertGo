
package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"store-manager/internal/domain/model"
	"store-manager/internal/domain/service"
)

type MockProductRepository struct{
	mock.Mock
}

func (m *MockProductRepository) FindByName(name string) (*model.Product, error){
	args := m.Called(name)
	if args.Get(0) == nil{
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductRepository) Save(product *model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func Test_DeberíaAlmacenarCorrectamente_AlExistirElProducto(t *testing.T){
	//Arrange (configuración)
	productName := "Camiseta"
	initialSotck := 10
	increment := 5

	product := &model.Product {
		Name: productName,
		Stock: initialSotck,
	}

	mockRepo := new(MockProductRepository)

	//Definción de los stubs (Expectativas para Go) para Buscar (FindByName) y Guardar (Save)
	mockRepo.On("FindByName", productName).
		Return(product, nil).
		Once()
	
	mockRepo.On("Save", mock.AnythingOfType("*model.Product")).
		Run(func(args mock.Arguments){
			p := args.Get(0).(*model.Product)
			assert.Equal(t, initialSotck + increment, p.Stock )
		}).
		Return(nil).
		Once()

	productService := service.NewProductService(mockRepo)

	//Act (ejecución)
	productService.IncrementStock(productName, increment)

	//Assert (validación)
	mockRepo.AssertExpectations(t)
}

func Test_DeberíaEnviarError_AlNoExistirElProducto(t *testing.T){
	//Arrange (configuración)
	productName := "No Existente"
	increment := 5

	mockRepo := new(MockProductRepository)

	//Definción de los stubs (Expectativas para Go) para Buscar (FindByName) y Guardar (Save)
	mockRepo.On("FindByName", productName).
		Return(nil, nil).
		Once()
	
	productService := service.NewProductService(mockRepo)

	//Act (ejecución)
	err := productService.IncrementStock(productName, increment)

	//Assert (validación)
	assert.Error(t, err)
	assert.Equal(t,"producto no encontrado",err.Error())
	mockRepo.AssertExpectations(t)
}