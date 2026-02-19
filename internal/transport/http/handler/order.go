// Package handler contains HTTP handlers.
//
// Handlers translate HTTP requests into service calls
// and map service responses to HTTP responses.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"botmanager/internal/domain"
	"botmanager/internal/service"
	"botmanager/internal/transport/http/dto"
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
		if errors.Is(err, domain.ErrProductNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.OrderReponse{
		ID:         order.ID(),
		CustomerID: order.CustomerID(),
		ProductID:  order.ProductID(),
		Price:      order.Price(),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

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
