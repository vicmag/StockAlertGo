package domain

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
    // Busca el producto por su ID
    product, err := s.repo.FindByID(productID)
    if err != nil {
        return err
    }

    // Actualiza el nivel mínimo de stock
    product.MinStock = minStock

    // Guarda el producto actualizado
    return s.repo.Save(product)
}