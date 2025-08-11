package handlers

import (
	"encoding/json"
	"net/http"
	"shopping-cart-app/src/services"
)

type HandlerStruct struct {
	productService *services.ServicesProdu
	serviceCart    *services.CartService
}

func NewHandler(productService *services.ServicesProdu, serviceCart *services.CartService) *HandlerStruct {
	return &HandlerStruct{productService: productService, serviceCart: serviceCart}
}

func (h *HandlerStruct) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *HandlerStruct) AddToCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		ProductID int `json:"product_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.serviceCart.AddProductToCar(requestBody.ProductID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cart := h.serviceCart.GetCart()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cart); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

}

func (h *HandlerStruct) GetCart(w http.ResponseWriter, r *http.Request) {
	cart := h.serviceCart.GetCart()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cart); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
