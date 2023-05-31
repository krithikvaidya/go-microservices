package app

import (
	"learning-http/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	r := mux.NewRouter()

	r.HandleFunc("/greet", handlers.Greet)
	r.HandleFunc("/customers", handlers.GetCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers", handlers.NewCustomer).Methods(http.MethodPost)

	log.Println("starting server ....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
