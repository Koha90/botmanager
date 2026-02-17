package domain

type DomainEvent interface {
	EventName() string
}

type OrderConfirmed struct {
	OrderID int
}

func (e OrderConfirmed) EventName() string {
	return "order_confirmed"
}
