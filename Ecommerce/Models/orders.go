package models

import "fmt"

type order struct {
	OrderID     string
	CustomerID  string
	TotalAmount float64
	Status      string
}

func NewOrder(orderID, customerID string, totalAmount float64) *order {
	return &order{
		OrderID:     orderID,
		CustomerID:  customerID,
		TotalAmount: totalAmount,
		Status:      "Pending",
	}
}
func (o *order) UpdateStatus(status string) {
	o.Status = status
}
func (o *order) IsCompleted() bool {
	return o.Status == "Completed"
}
func (o *order) IsPending() bool {
	return o.Status == "Pending"
}
func (o *order) Summary() string {
	return "Order " + o.OrderID + " for Customer " + o.CustomerID + " is " + o.Status + " with total amount $" + fmt.Sprintf("%.2f", o.TotalAmount)
}
