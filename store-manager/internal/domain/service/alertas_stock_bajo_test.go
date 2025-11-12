package service_test

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"store-manager/internal/domain/model"
	"store-manager/internal/domain/service"
)

type MockProductoDB struct{
	mock.Mock
}

func (m *MockProductoDB) FindByName(nombre string) (*model.Producto, error){
	args := m.Called(nombre)
	return args.Get(0).(*model.Producto), args.Error(1)
}

func (m *MockProductoDB) Save(producto *model.Producto) error{
	args := m.Called(producto)
	return args.Error(0)
}

func TestIncrementoStock_DeberiaGuardarActualización_AlIncrementarStock(t *testing.T){
	//Arrange (configuración)
	nombreProducto := "Camiseta"
	stockInicial := 10
	incremento := 5

	producto := &model.Producto{
		Nombre: nombreProducto,
		Stock: stockInicial,
	}

	//Mock de BD (repositorio)
	mockDB := new(MockProductoDB)

	//Configuración del comportamiento de la BD (expectativas/stubs)
	mockDB.On("FindByName", nombreProducto).
		Return(producto,nil).
		Once()

	mockDB.On("Save", mock.AnythingOfType("*model.Producto")).
		Run(func(args mock.Arguments){
			p := args.Get(0).(*model.Producto)
			assert.Equal(t, stockInicial+incremento, p.Stock)
		}).
		Return(nil).
		Once()

	productoService := service.NewProductService(mockDB)

	//Act (ejecución)
	err := productoService.IncrementaStock(nombreProducto, incremento)

	//Assert (validación)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}