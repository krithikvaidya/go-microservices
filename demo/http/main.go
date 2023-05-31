package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Id      string `json:"id"`
	Name    string `json:"full_name"`
	City    string `json:"city,omitempty"`
	Zipcode string `json:"-"`
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"1", "Abhay", "Blr", "123456"},
		{"2", "Ashish", "", "110011"},
		{"3", "Rob", "Mum", "220011"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
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
