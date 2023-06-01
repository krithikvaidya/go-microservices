package handlers

import (
	"encoding/json"
	"learning-http/service"
	"net/http"
)

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	service := service.NewCustomerService()

	customers, err := service.GetAllCustomers()
	if err != nil {
		// we have some error here
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
