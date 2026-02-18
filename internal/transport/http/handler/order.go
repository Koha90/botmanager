// Package handler contains HTTP handlers.
//
// Handlers translate HTTP requests into service calls
// and map service responses to HTTP responses.
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"botmanager/internal/service"
)

// OrderHandler handles HTTP requests related to orders.
//
// It does not contain business logic.
// It only coordinates requests parsing and response formatting.
type OrderHandler struct {
	service *service.OrderService
}

// NewOrderHandler ...
func NewOrderHandler(s *service.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

type createRequest struct {
	CustomerID int `json:"customer_id"`
	ProductID  int `json:"product_id"`
}

// Create handles order creation request.
//
// Expects JSON body:
//
//	{
//	  "customer_id": int,
//		"product_id": int
//	}
//
// Returns created order as JSON.
func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.service.Create(r.Context(), req.CustomerID, req.ProductID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// Confirm handles order confirmation.
//
// Path param:
//
//	id - order indentifier
//
// Returns 204 No Content on success.
func (h *OrderHandler) Confirm(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := h.service.Confirm(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Cancel handles order cancellation.
//
// Path param:
//
//	id - order indentifier
//
// Returns 204 No Content on success.
func (h *OrderHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	if err := h.service.Cancel(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
