// Package dto ...
package dto

type OrderReponse struct {
	ID         int   `json:"id"`
	CustomerID int   `json:"customer_id"`
	ProductID  int   `json:"product_id"`
	Price      int64 `json:"price"`
}
