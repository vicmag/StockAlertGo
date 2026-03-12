package services_test

import(
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

func (m *MockProductRepository) Save(product *models.Product) error{
	args := m.Called(product)
	return args.Error(0)
}

func TestIncrementStock_WhenProductExists_ShouldIncrementStock(t *testing.T){
	//Arrange (configuración)
	productName := "Camiseta"
	initialStock := 10
	increment := 5
	finalStock := initialStock + increment

	product := &models.Product{
		Name: productName,
		Stock: initialStock,
	}

	mockRepo := new(MockProductRepository)
	//Configuración de las "expectativas"
	mockRepo.On("FindByName", productName).
		Return(product, nil).
		Once()

	mockRepo.On("Save", mock.MatchedBy(func(p *models.Product) bool{
		return p.Stock == finalStock
		})).
		Return(nil).
		Once()

	service := services.NewProductService(mockRepo)
	
	//Act (ejecución)
	res := service.IncrementStock(productName, increment)

	//Assert (validación)
	assert.NoError(t, res)
	mockRepo.AssertExpectations(t)
}