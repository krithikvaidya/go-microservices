package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", greet)

	log.Println("starting server 1....")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("{ \"msg\": \"%s\"}", "Hello World!!")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(msg))
}
