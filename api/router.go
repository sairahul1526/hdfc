package api

import (
	"encoding/json"
	"net/http"

	UserAPI "hdfc-backend/api/user"

	"github.com/gorilla/mux"
)

// HealthCheck .
// for load balancer/beanstalk to know whether server/ec2 is healthy
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ok")
}

// LoadRouter - get mux router with all the routes
func LoadRouter() *mux.Router {
	router := mux.NewRouter()

	// middlewares
	router.Use(JSONHeaderMiddleware)

	UserAPI.LoadUserRoutes(router)

	router.Path("/").HandlerFunc(HealthCheck).Methods("GET")

	return router
}
