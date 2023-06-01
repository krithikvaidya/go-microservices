package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type Customer struct {
	Id      string `json:"id" xml:"id"`
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city,omitempty" xml:"city,omitempty"`
	Zipcode string `json:"-" xml:"zip"`
}

func NewCustomer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("New customer added"))
}

// localhost:8080/customers/123
// localhost:8080/customers/abc -> error/404

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	w.Write([]byte(customerId))
}

func GetCustomers(c *gin.Context) {
	customers := []Customer{
		{"1", "Abhay", "Blr", "123456"},
		{"2", "Ashish", "", "110011"},
		{"3", "Rob", "Mum", "220011"},
	}
	c.JSON(http.StatusOK, customers)
	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// }
	// if r.Header.Get("Content-Type") == "application/json" {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }
	// if r.Header.Get("Content-Type") == "" {
	// 	w.WriteHeader(404)
	// }
}

func Greet(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	msg := fmt.Sprintf("{ \"msg\": \"ID is: %s\"}", id)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(msg))
}
