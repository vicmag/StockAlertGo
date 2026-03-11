package services_test

import(
	"testing"
	"github.com/stretchr/testify/mock"
)


func TestIncrementStock_WhenProductExists_ShouldIncrementStock(t *testing.T){
	//Arrange (configuración)
	productName := "Camiseta"
	initialStock := 10
	increment := 5

	product := &models.Product{
		Name: productName,
		Stock: initialStock.
	}

	
	//Act (ejecución)
	//Assert (validación)
}