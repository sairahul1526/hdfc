package product

import "github.com/gorilla/mux"

// LoadProductRoutes - load all product routes
func LoadProductRoutes(router *mux.Router) {
	router.HandleFunc("/categories", CategoriesGet).Methods("GET")
	router.HandleFunc("/products", ProductsGet).Queries(
		"category_id", "{category_id}",
	).Methods("GET")
	router.HandleFunc("/product", ProductGet).Queries(
		"id", "{id}",
	).Methods("GET")
}
