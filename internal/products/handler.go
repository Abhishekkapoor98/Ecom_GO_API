package products

import (
	"log"
	"net/http"

	"github.com/Abhishekkapoor98/Ecom_GO_API/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	// 1. Call the service to list the products.
	// Service -> ListProducts
	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}
