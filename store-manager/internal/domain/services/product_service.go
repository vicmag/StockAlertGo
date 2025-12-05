package services

import (
    "store-manager/internal/domain/ports"
    "store-manager/internal/events"
    "store-manager/internal/observers"
)

type ProductService struct {
    repo      ports.ProductRepository
    notifier  ports.NotifierService
    observers []observers.StockObserver
}

// NewProductServiceWithoutNotifier - Constructor para tests sin notificación
func NewProductServiceWithoutNotifier(repo ports.ProductRepository) *ProductService {
    return &ProductService{
        repo:      repo,
        notifier:  nil,
        observers: []observers.StockObserver{},
    }
}

// NewProductService - Constructor con notificador (para tests con notificación)
func NewProductService(repo ports.ProductRepository, notifier ports.NotifierService) *ProductService {
    service := &ProductService{
        repo:      repo,
        notifier:  notifier,
        observers: []observers.StockObserver{},
    }
    
    // Configurar el observer de stock bajo si hay notificador
    if notifier != nil {
        service.setupLowStockObserver()
    }
    
    return service
}

// DecrementStock - Implementación refactorizada
func (s *ProductService) DecrementStock(productName string, amount int) error {
    // 1. Buscar producto
    product, _ := s.repo.FindByName(productName)
    
    // 2. Preservar stock anterior
    oldStock := product.Stock
    
    // 3. Decrementar stock
    product.Stock -= amount
    
    // 4. Guardar cambios
    _ = s.repo.Save(product)
    
    // 5. Crear y notificar evento
    event := events.StockChangeEvent{
        Product:   product,
        OldStock:  oldStock,
        NewStock:  product.Stock,
        Operation: "decrement",
    }
    
    s.notifyObservers(event)
    
    return nil
}

// MÉTODOS INTERNOS PARA OBSERVER PATTERN

// setupLowStockObserver - Configura observer para notificaciones de stock bajo
func (s *ProductService) setupLowStockObserver() {
    lowStockObserver := &lowStockObserver{notifier: s.notifier}
    s.observers = append(s.observers, lowStockObserver)
}

// notifyObservers - Notifica a todos los observadores
func (s *ProductService) notifyObservers(event events.StockChangeEvent) {
    for _, observer := range s.observers {
        observer.OnStockChanged(event)
    }
}

// Subscribe - Para extensiones futuras (pública)
func (s *ProductService) Subscribe(observer observers.StockObserver) {
    s.observers = append(s.observers, observer)
}

// lowStockObserver - Implementación interna
type lowStockObserver struct {
    notifier ports.NotifierService
}

func (l *lowStockObserver) OnStockChanged(event events.StockChangeEvent) {
    if l.shouldNotifyLowStock(event) {
        l.notifier.SendLowStockAlert(
            event.Product.Name,
            event.NewStock,
            event.Product.MinStockLevel,
        )
    }
}

func (l *lowStockObserver) shouldNotifyLowStock(event events.StockChangeEvent) bool {
    product := event.Product
    return event.NewStock < product.MinStockLevel && 
           event.OldStock >= product.MinStockLevel
}