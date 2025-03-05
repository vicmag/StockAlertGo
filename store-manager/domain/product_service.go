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
    product, err := s.repo.FindByID(productID)
    if err != nil {
        return err
    }
    product.MinStock = minStock
    return s.repo.Save(product)
}
    