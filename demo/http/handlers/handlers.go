package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Id      string `json:"id" xml:"id"`
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city,omitempty" xml:"city,omitempty"`
	Zipcode string `json:"-" xml:"zip"`
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"1", "Abhay", "Blr", "123456"},
		{"2", "Ashish", "", "110011"},
		{"3", "Rob", "Mum", "220011"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
	if r.Header.Get("Content-Type") == "" {
		w.WriteHeader(404)
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("{ \"msg\": \"%s\"}", "Hello World!!")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(msg))
}
