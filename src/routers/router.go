package routers

import (
	"net/http"
	"shopping-cart-app/src/handlers"
	"shopping-cart-app/src/services"
)

func Router(mux *http.ServeMux) {
	serviceProduct := services.NewServiceCar()
	cartService := services.NewCartService(serviceProduct)
	handler := handlers.NewHandler(serviceProduct, cartService)

	mux.HandleFunc("/api/products", handler.GetProductsHandler)
	mux.HandleFunc("GET /api/cart", handler.GetCart)
	mux.HandleFunc("POST /api/cart", handler.AddToCarHandler)

}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
