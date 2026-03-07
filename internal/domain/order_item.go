package domain

// OrderItem represents snapshot of product at purchase time.
type OrderItem struct {
	productID int
	variantID int
	quantity  int
	unitPrice int64
}

// NewOrderItem creates a new order item instance.
func NewOrderItem(
	productID int,
	variantID int,
	quantity int,
	unitPrice int64,
) OrderItem {
	return OrderItem{
		productID: productID,
		variantID: variantID,
		quantity:  quantity,
		unitPrice: unitPrice,
	}
}

// ProductID returns product id.
func (i OrderItem) ProductID() int {
	return i.productID
}

// VariantID returns variant id.
func (i OrderItem) VariantID() int {
	return i.variantID
}

// Quantity returns quantity of a quantity item.
func (i OrderItem) Quantity() int {
	return i.quantity
}

// UnitPrice returns one unit price.
func (i OrderItem) UnitPrice() int64 {
	return i.unitPrice
}

// Total returns total price.
func (i OrderItem) Total() int64 {
	return int64(i.quantity) * i.unitPrice
}
