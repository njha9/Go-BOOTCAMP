package models

import "fmt"

type product struct {
	ProductID   string
	Name        string
	Description string
	Price       float64
	Stock       int
}

func NewProduct(productID, name, description string, price float64, stock int) *product {
	return &product{
		ProductID:   productID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
}

func (p *product) IsInStock() bool {
	return p.Stock > 0
}

func (p *product) UpdateStock(quantity int) {
	p.Stock += quantity
}

func (p *product) Summary() string {
	return p.Name + ": " + p.Description + " - $" + fmt.Sprintf("%.2f", p.Price) + " (Stock: " + fmt.Sprintf("%d", p.Stock) + ")"
}
