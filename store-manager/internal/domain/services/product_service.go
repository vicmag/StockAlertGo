package services

import "store-manager/internal/domain/ports"

// ProductService - Estructura mínima
type ProductService struct {
    repo ports.ProductRepository
    notifier  ports.NotifierService // ← NUEVO: Servicio de notificación
}


// NewProductService - Constructor principal actualizado
func NewProductService(repo ports.ProductRepository, notifier ports.NotifierService) *ProductService {
    return &ProductService{
        repo:     repo,
        notifier: notifier,
    }
}

// NewProductServiceWithoutNotifier - Constructor para compatibilidad
func NewProductServiceWithoutNotifier(repo ports.ProductRepository) *ProductService {
    return &ProductService{
        repo:     repo,
        notifier: nil, // ← Notificador nulo para compatibilidad
    }
}


func (s *ProductService) DecrementStock(productName string, amount int) error {
    // 1. Buscar producto por nombre
    product, _ := s.repo.FindByName(productName)
   
    // 2. Decrementar el stock
    product.Stock -= amount
   
    // 3. NUEVO: Verificar si se debe enviar alerta de stock bajo
    if s.notifier != nil && product.Stock < product.MinStockLevel {
        // Notificar si el stock está por debajo del mínimo
        s.notifier.SendLowStockAlert(product.Name, product.Stock, product.MinStockLevel)
    }
   
    // 4. Guardar el producto actualizado
    _ = s.repo.Save(product)
   
    return nil
}