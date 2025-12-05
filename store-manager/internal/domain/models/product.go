package models

type Product struct{
	Name string
	Stock int
	MinStockLevel int // ← NUEVO: Nivel mínimo para alertas
}