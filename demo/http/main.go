package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Id      string
	Name    string
	City    string
	Zipcode string
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getCustomers)

	log.Println("starting server 1....")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("{ \"msg\": \"%s\"}", "Hello World!!")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(msg))
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"1", "Abhay", "Blr", "123456"},
		{"2", "Ashish", "Del", "110011"},
		{"3", "Rob", "Mum", "220011"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
