package app

import (
	"learning-http/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {

	// r := mux.NewRouter()

	// r.HandleFunc("/greet", handlers.Greet)
	// r.HandleFunc("/customers", handlers.GetCustomers).Methods(http.MethodGet)
	// r.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.GetCustomer).Methods(http.MethodGet)
	// r.HandleFunc("/customers", handlers.NewCustomer).Methods(http.MethodPost)

	// log.Println("starting server ....")
	// log.Fatal(http.ListenAndServe("localhost:8080", r))
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/customers", handlers.GetCustomers)
	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
