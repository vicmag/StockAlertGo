package domain

// Product representa un producto en el inventario
type Product struct {
    ID           string
    Name         string
    CurrentStock int
    MinStock     int
}