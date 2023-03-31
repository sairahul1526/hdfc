package user

import (
	"github.com/gorilla/mux"

	OrderAPI "hdfc-backend/api/user/order"
	ProductAPI "hdfc-backend/api/user/product"
)

// LoadUserRoutes - load all user routes with user prefix
func LoadUserRoutes(router *mux.Router) {
	userRoutes := router.PathPrefix("/user").Subrouter()

	ProductAPI.LoadProductRoutes(userRoutes)
	OrderAPI.LoadProductRoutes(userRoutes)
}
