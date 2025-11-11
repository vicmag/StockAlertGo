package service_test

import (
	"testing"
	"github.com/stretchr/testify/mock"
)

type MockProductoDB struct{
	mock.Mock
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
		Return(producto,nil)

	//Act (ejecución)

	//Assert (validación)
}