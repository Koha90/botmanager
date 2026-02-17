// Package domain provide struct of models.
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

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int64
	ImageURL    *string
	DistictID   int
	CategoryID  int
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
