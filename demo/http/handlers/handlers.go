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
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) CustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	// service := service.NewCustomerService()

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		// we have some error here
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func NewCustomerHandler(svc service.CustomerService) CustomerHandlers {
	return CustomerHandlers{svc}
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
