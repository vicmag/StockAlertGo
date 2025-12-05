package ports

// NotifierService - Interfaz simple para notificaciones
type NotifierService interface {
    SendLowStockAlert(productName string, currentStock, minStockLevel int) error
}