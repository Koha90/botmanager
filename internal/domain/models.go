// Package domain contains the core business model of the application.
//
// It defines aggregates, value objects, domain events and business invariants.
// The domain layer does not depend on infrastructure, transport or persistance.
//
// All business rules must live here.
package domain

import "time"

type City struct {
	ID   int
	Name string
}

type Category struct {
	ID     string
	Name   string
	CityID int
}

type District struct {
	ID         int
	Name       string
	CategoryID int
	CityID     int
}

type Customer struct {
	ID         int
	TelegramID int64
	Username   string
	Balance    int64
}

type Transaction struct {
	ID         int
	CustomerID int
	Amount     int64
	Reason     string
	CreatedAt  time.Time
}
