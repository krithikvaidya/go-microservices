package handlers

import (
	"encoding/json"
	"learning-http/service"
	"net/http"
)

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	service := service.NewCustomerService()

	customers := service.GetAllCustomers()
	if customers != nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(404)
	}
}
