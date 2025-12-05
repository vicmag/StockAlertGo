package events

import "store-manager/internal/domain/models"

type StockChangeEvent struct {
    Product     *models.Product
    OldStock    int
    NewStock    int
    Operation   string
}