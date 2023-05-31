package app

import (
	"learning-http/handlers"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", handlers.Greet)
	http.HandleFunc("/customers", handlers.GetCustomers)

	log.Println("starting server 1....")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
