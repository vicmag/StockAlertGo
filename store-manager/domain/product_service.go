package domain

import (
    "errors" // Importación del paquete errors
)

// ProductService contiene la lógica de negocio para gestionar productos
type ProductService struct {
    repo ProductRepository
}

// NewProductService crea una nueva instancia de ProductService
func NewProductService(repo ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

// SetMinStock establece el nivel mínimo de stock para un producto
func (s *ProductService) SetMinStock(productID string, minStock int) error {
    if minStock <= 0 {
        return ErrInvalidMinStock // Refactor: Validación adicional para mejorar la robustez
    }

    product, err := s.repo.FindByID(productID)
    if err != nil {
        return err
    }

    product.MinStock = minStock
    return s.repo.Save(product)
}

// CheckLowStock verifica si el stock de un producto está por debajo del nivel mínimo
func (s *ProductService) CheckLowStock(productID string) bool {
    // Implementación vacía para que la prueba falle
    return true
}

// ErrInvalidMinStock es un error personalizado para niveles mínimos inválidos
var ErrInvalidMinStock = errors.New("el nivel mínimo de stock debe ser mayor que cero")
    