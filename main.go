package main

import (
	"log"
	"net/http"
	"shopping-cart-app/src/routers"
)

func main() {

	mux := http.NewServeMux()
	routers.Router(mux)

	handlerC := routers.CorsMiddleware(mux)

	log.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", handlerC)
}
