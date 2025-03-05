package domain

// ProductRepository define la interfaz para interactuar con el almacenamiento de productos
type ProductRepository interface {
    FindByID(id string) (*Product, error)
    Save(product *Product) error
}