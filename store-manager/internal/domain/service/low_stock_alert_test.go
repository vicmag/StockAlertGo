import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
)


func TestIncrementStock_DeberiaGuardarElProductoCorrectamente_AlIncrementar(t *Testing.T){
	//Arrange (Configuración)
	productName := "Camiseta"
	initialStock := 10
	increment := 5
	finalStock := 15

	product := &model.Product {
		Name: productName,
		Stock: initialStock,
	}

	//Act (Ejcución)
	//Assert (Validación)

}