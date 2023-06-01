package handlers

import (
	"encoding/json"
	"learning-http/service"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) CustomersHandler(w http.ResponseWriter, r *http.Request) {
	// service := service.NewCustomerService()

	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		// we have some error here
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) CustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	// service := service.NewCustomerService()

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		// we have some error here
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

func NewCustomerHandler(svc service.CustomerService) CustomerHandlers {
	return CustomerHandlers{svc}
}
