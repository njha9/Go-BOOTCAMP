package models

type CartItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

type Cart struct {
	Items []CartItem
}

func (c *Cart) AddItem(item CartItem) {
	c.Items = append(c.Items, item)
}

func (c *Cart) TotalAmount() float64 {
	var total float64
	for _, item := range c.Items {
		total += float64(item.Quantity) * item.Price
	}
	return total
}
func (c *Cart) ItemCount() int {
	return len(c.Items)
}
func (c *Cart) Clear() {
	c.Items = []CartItem{}
}
func NewCart() *Cart {
	return &Cart{
		Items: []CartItem{},
	}
}
