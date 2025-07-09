package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"store-manager/internal/domain/model"
	"store-manager/internal/domain/service"
)


type testContext struct{
	mockRepo *MockProductRepository
	productService *service.ProductService
}

func setup() testContext{
	mockRepo := new(MockProductRepository)
	productService := service.NewProductService(mockRepo)

	return testContext{mockRepo: mockRepo, productService: productService}
}

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

func (m *MockProductRepository) Save(product *model.Product) error{
	args := m.Called(product)
	return args.Error(0)
}

func TestIncrementStockCases(t *testing.T) {
	t.Run("debería almacenarse correctamente al existir el producto", func(t *testing.T) {
		//Estructura (Patrón) AAA
		//Arrange (configuración)
		ctx := setup()
		productName := "Producto Inexistente"
		initialStock := 10
		increment := 5

		product := &model.Product{
			Name: productName,
			Stock: initialStock,
		}

		//Configuración del expectativos del mock
		ctx.mockRepo.On("FindByName", productName).
			Return(product,nil).
			Once()

			ctx.mockRepo.On("Save",mock.AnythingOfType("*model.Product")).
			Run(func(args mock.Arguments) { 
				p := args.Get(0).(*model.Product)
				assert.Equal(t, initialStock+increment, p.Stock)
			}).
			Return(nil).
			Once()
		

		//Act (ejecución)
		ctx.productService.IncrementStock(productName, increment)

		//Assert (validación)
		ctx.mockRepo.AssertExpectations(t)
	})

	t.Run("debería enviar error al no existir el producto", func(t *testing.T) {
		//Estructura (Patrón) AAA
		//Arrange (configuración)
		ctx := setup()
		productName := "Camiseta"
		increment := 5

		//Configuración del expectativos del mock
		ctx.mockRepo.On("FindByName", productName).
			Return(nil,nil).
			Once()


		//Act (ejecución)
		err := ctx.productService.IncrementStock(productName, increment)

		//Assert (validación)
		assert.Error(t, err)
		assert.Equal(t, "producto no encontrado", err.Error())
		ctx.mockRepo.AssertNotCalled(t, "Save", mock.Anything)
		ctx.mockRepo.AssertExpectations(t)
	})
}

