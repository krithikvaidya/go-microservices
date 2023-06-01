package app

import (
	"learning-http/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	r := mux.NewRouter()

	r.HandleFunc("/customers", handlers.CustomerHandler).Methods(http.MethodGet)

	log.Println("starting server ....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
