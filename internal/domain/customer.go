package domain

// Customer represent the customer.
type Customer struct {
	id         int
	telegramID int64
	username   string
	balance    int64
}

// NewCustomer creates new custromer.
func NewCustomer(telegramID int64, username string) *Customer {
	return &Customer{
		telegramID: telegramID,
		username:   username,
		balance:    0,
	}
}

// TelegramID returns telegram identifier of customer.
func (c *Customer) TelegramID() int64 {
	return c.telegramID
}

// ID returns identifier customer in application.
func (c *Customer) ID() int {
	return c.id
}
