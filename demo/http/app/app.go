package app

import (
	"learning-http/handlers"
	"learning-http/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	r := mux.NewRouter()
	svc := service.NewCustomerService()

	// application wiring
	ch := handlers.NewCustomerHandler(svc)

	r.HandleFunc("/customers", ch.CustomersHandler).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id}", ch.CustomerHandler).Methods(http.MethodGet)

	log.Println("starting server ....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
