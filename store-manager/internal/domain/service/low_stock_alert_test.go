package service_test
import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"store-manager/internal/domain/model"
	"store-manager/internal/domain/service"
)

type MockProductRepository struct{
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

func TestIncrementStock_DeberiaGuardarElProductoCorrectamente_AlIncrementar(t *testing.T){
	//Arrange (Configuración)
	productName := "Camiseta"
	initialStock := 10
	increment := 5

	product := &model.Product {
		Name: productName,
		Stock: initialStock,
	}

	//Base de datos
	mockDB := new(MockProductRepository);

	//Defino el comportamiento de mi BD
	//Expectativas (stubs)
	mockDB.On("FindByName", productName).
		Return(product, nil).
		Once()

	mockDB.On("Save", mock.AnythingOfType("*model.Product")).
		Run(func(args mock.Arguments){
			p := args.Get(0).(*model.Product)
			assert.Equal(t, initialStock+increment, p.Stock) 
		}).
		Return(nil).
		Once()

	productService := service.NewProductService(mockDB) 

	//Act (Ejcución)
	productService.IncrementStock(productName, increment)

	//Assert (Validación)
	mockDB.AssertExpectations(t)

}