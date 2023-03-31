package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// StartServer - start server using mux
func StartServer() {
	// ec2 router
	fmt.Println(http.ListenAndServe(":"+os.Getenv("port"), &WithCORS{LoadRouter()}))
}

func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// cors configuration
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	res.Header().Set("Access-Control-Max-Age", "86400")

	if req.Method == "OPTIONS" {
		return
	}

	s.r.ServeHTTP(res, req)
}

// WithCORS .
type WithCORS struct {
	r *mux.Router
}
