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
	productName := "Camiseta"
	initialStock := 10
	increment := 5

	product := &model.Product{
		Name:   productName,
		Stock:  initialStock
	}

	//Configuro las expectativas del mock
	mockRepo.On("FindByName", productName).
		Return(product, nil).
		Once()


	 
	// Act (ejecución)

	// Assert (validación)

}