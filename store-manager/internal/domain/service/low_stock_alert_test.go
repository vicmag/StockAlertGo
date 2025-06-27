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

type testContext struct {
	mockRepo *MockProductRepository
	productService *service.ProductService
}

func setup() testContext {
	mockRepo := new(MockProductRepository)
	productService := service.NewProductService(mockRepo)
	return testContext{
		mockRepo: mockRepo,
		productService: productService}
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

func TestCuandoIncrementoElStockDelProductoEntoncesSeAlmacenaCorrectamente(t *testing.T){
	//Arrange (Configuración)
	ctx := setup()

	productName := "Camiseta"
	initialStock := 10
	increment := 5

	product :=  &model.Product{
		Name:  productName,
		Stock: initialStock,
	}

	//Configuración de expectativas
	ctx.mockRepo.On("FindByName", productName).
		Return(product, nil).
		Once()
	
	ctx.mockRepo.On("Save", product).
		Run(
			func(args mock.Arguments){
				p := args.Get(0).(*model.Product)
				assert.Equal(t, initialStock+increment, p.Stock)
		}).
		Return(nil).
		Once()


	//Act (Ejecución)
	ctx.productService.IncrementStock(productName, increment)

	//Assert (Verificación)
	ctx.mockRepo.AssertExpectations(t)
}

func TestCuandoIncrementoElStockDeUnProductoInexistent_EntoncesObtengoUnError(t *testing.T){
	// Arrange (Configuración)
	ctx := setup()

	productName := "Camiseta"	
	increment := 5

	//Configuración de expectativas
	ctx.mockRepo.On("FindByName", productName).
		Return(nil, nil).
		Once()

	// Act (Ejecución)
	err := ctx.productService.IncrementStock(productName, increment)

	// Assert (Verificación)
	assert.Error(t, err)
	assert.Equal(t, "producto no encontrado", err.Error())
	ctx.mockRepo.AssertNotCalled(t, "Save", mock.Anything)
	ctx.mockRepo.AssertExpectations(t)

}

func TestCuandoElStockEsNegativoEntoncesError(t *testing.T) {
	//Arrange (Configuración)
	ctx := setup()

	productName := "Camiseta"	
	increment := -5

	// Act (Ejecución)
	err := ctx.productService.IncrementStock(productName, increment)

	// Assert (Verificación)}
	assert.Error(t, err)
	assert.Equal(t, "El incremento  ser positivo", err.Error())

	ctx.mockRepo.AssertNotCalled(t, "FindByName", mock.Anything)
	ctx.mockRepo.AssertNotCalled(t, "Save", mock.Anything)

}

func TestCuandoElStockEsCeroEntoncesError(t *testing.T) {
	//Arrange (Configuración)
	ctx := setup()

	productName := "Camiseta"	
	increment := 0

	// Act (Ejecución)
	err := ctx.productService.IncrementStock(productName, increment)

	// Assert (Verificación)}
	assert.Error(t, err)
	assert.Equal(t, "El incremento  ser positivo", err.Error())

	ctx.mockRepo.AssertNotCalled(t, "FindByName", mock.Anything)
	ctx.mockRepo.AssertNotCalled(t, "Save", mock.Anything)

}
