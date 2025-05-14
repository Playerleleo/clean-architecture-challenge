package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/usecase"
)

type OrderHandler struct {
	ListOrdersUseCase *usecase.ListOrdersUseCase
}

func NewOrderHandler(listOrdersUseCase *usecase.ListOrdersUseCase) *OrderHandler {
	return &OrderHandler{
		ListOrdersUseCase: listOrdersUseCase,
	}
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	orders, err := h.ListOrdersUseCase.Execute(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
