package app

import (
	"learning-http/domain"
	"learning-http/handlers"
	"learning-http/logger"
	"learning-http/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	r := mux.NewRouter()

	// application wiring
	dbConn := getDbClient()
	customerRepo := domain.NewCustomerRepositoryDb(dbConn)
	// stubRepo := domain.NewStubCustomerRepository()

	svc := service.NewCustomerService(&customerRepo)
	ch := handlers.NewCustomerHandler(svc)

	r.HandleFunc("/customers", ch.CustomersHandler).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id}", ch.CustomerHandler).Methods(http.MethodGet)

	// log.Println("starting server ....")
	logger.Info("starting server ....")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getDbClient() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:student@tcp(localhost:3307)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
