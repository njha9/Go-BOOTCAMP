package main

import "fmt"

type CustomerRating struct {
	CustomerID string
	Rating     int
	Comment    string
}

func NewCustomerRating(customerID string, rating int, comment string) *CustomerRating {
	return &CustomerRating{
		CustomerID: customerID,
		Rating:     rating,
		Comment:    comment,
	}
}
func (cr *CustomerRating) IsPositive() bool {
	return cr.Rating >= 4
}

func (cr *CustomerRating) IsNegative() bool {
	return cr.Rating <= 2
}

func (cr *CustomerRating) Summary() string {
	return fmt.Sprintf("Customer %s rated %d: %s", cr.CustomerID, cr.Rating, cr.Comment)
}

func main() {
	rating := NewCustomerRating("C123", 5, "Excellent service!")
	fmt.Println(rating.Summary())
	fmt.Println("Is Positive:", rating.IsPositive())
	fmt.Println("Is Negative:", rating.IsNegative())
}
