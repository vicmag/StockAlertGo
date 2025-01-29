package model

type Product struct {
	ID                string
	Name              string
	CurrentStock      int
	LowStockThreshold int
}
