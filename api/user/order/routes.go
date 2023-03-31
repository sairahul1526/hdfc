package product

import "github.com/gorilla/mux"

// LoadProductRoutes - load all product routes
func LoadProductRoutes(router *mux.Router) {
	router.HandleFunc("/orders", OrdersGet).Queries(
		"customer_id", "{customer_id}",
	).Methods("GET")
	router.HandleFunc("/order", OrderPost).Methods("POST")
	router.HandleFunc("/order", OrderUpdate).Queries(
		"id", "{id}",
	).Methods("PATCH")
}
