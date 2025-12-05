package observers

import "store-manager/internal/events"

// StockObserver - Interfaz para el patrón Observer
type StockObserver interface {
    OnStockChanged(event events.StockChangeEvent)
}