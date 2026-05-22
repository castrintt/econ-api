package products

import (
	"log/slog"
	"net/http"

	"github.com/castrintt/econ-api/internal/shared"
)

// structs
type handler struct {
	service Service
}

// constructor
func NewHandler(service Service) *handler {
	return &handler{service: service}
}

// methods
func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	context := r.Context()
	// limit := r.URL.Query().Get("limit")
	// offset := r.URL.Query().Get("offset")
	// name := r.URL.Query().Get("name")

	products, err := h.service.GetProducts(context)
	if err != nil {
		slog.Error("Error getting products", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shared.WriteAsJSON(w, http.StatusOK, products)
}
