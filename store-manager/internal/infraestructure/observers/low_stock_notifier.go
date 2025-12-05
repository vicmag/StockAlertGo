package observers

import (
    "store-manager/internal/events"
    "store-manager/internal/domain/ports"
)

type LowStockNotifier struct {
    notifier ports.NotifierService
}

func NewLowStockNotifier(notifier ports.NotifierService) *LowStockNotifier {
    return &LowStockNotifier{
        notifier: notifier,
    }
}

func (l *LowStockNotifier) OnStockChanged(event events.StockChangeEvent) {
    if l.shouldNotifyLowStock(event) {
        l.notifier.SendLowStockAlert(
            event.Product.Name,
            event.NewStock,
            event.Product.MinStockLevel,
        )
    }
}

func (l *LowStockNotifier) shouldNotifyLowStock(event events.StockChangeEvent) bool {
    product := event.Product
    return event.NewStock < product.MinStockLevel && 
           event.OldStock >= product.MinStockLevel
}