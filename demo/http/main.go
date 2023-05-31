package main

import (
	"log"
	"net/http"
)

func main() {
	// server 1
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet2)

	go func() {
		log.Println("starting server 1....")
		log.Fatal(http.ListenAndServe("localhost:8081", mux))
	}()

	// server 2
	http.HandleFunc("/greet", greet)
	log.Println("starting server 2....")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!"))
}

func greet2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!! from new mux server"))
}
