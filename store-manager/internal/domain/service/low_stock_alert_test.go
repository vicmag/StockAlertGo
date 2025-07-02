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
	if args.Get(0) == nil {
		return nil, args.Error(1)	
	}
	return args.Get(0).(*model.Product), args.Error(1)
}


func (m *MockProductRepository) Save(product *model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}


//func TestCuandoIncrementoStock_EntoncesSeAlmacena
//func TestIncrementStock_IncremtoStock_ProductoSeAlmacena
func TestIncrementStock_DeberíaGuardarProductoConStockActualizado(t *testing.T){
	//Arrange (configuración)
	productName := "Camiseta"
	initialStock := 10
	increment := 5

	product := &model.Product{
		Name: productName,
		Stock: initialStock,
	}

	//Repositorio
	mockRepo := new(MockProductRepository);

	//Servicio
	productService := service.NewProductService(mockRepo)

	//Configuración de expectativas del mock
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

	//Act (ejecución)
	productService.IncrementStock(productName, increment)

	//Assert (validación)
	mockRepo.AssertExpectations(t)
}

func TestIncrementStock_DeberíaGenerarErrorConProductoInexistente(t *testing.T){
	//Arrange (configuración)
	productName := "Producto Inexistente"
	increment := 5

	//Repositorio
	mockRepo := new(MockProductRepository);

	//Servicio
	productService := service.NewProductService(mockRepo)

	//Configuración de expectativas del mock
	mockRepo.On("FindByName", productName).
		Return(nil, nil).
		Once()

	//Act (ejecución)
	err := productService.IncrementStock(productName, increment)

	//Assert (validación)
	assert.Error(t, err)
	assert.Equal(t, "Producto no encontrado", err.Error())
	mockRepo.AssertExpectations(t)
}