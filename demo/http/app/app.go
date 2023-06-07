package app

import (
	"fmt"
	"learning-http/domain"
	"learning-http/handlers"
	"learning-http/logger"
	"learning-http/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {

	sanityCheck()

	r := mux.NewRouter()

	// application wiring
	dbConn := getDbClient()
	customerRepo := domain.NewCustomerRepositoryDb(dbConn)
	// stubRepo := domain.NewStubCustomerRepository()

	svc := service.NewCustomerService(&customerRepo)
	ch := handlers.NewCustomerHandler(svc)
	am := handlers.NewAuthMiddleware()

	r.HandleFunc("/customers", ch.CustomersHandler).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id}", ch.CustomerHandler).Methods(http.MethodGet)

	r.Use(am.AuthMiddlewareHandler)
	r.Use(loggingMiddleware)

	logger.Info("starting server ....")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		logger.Info(fmt.Sprintf("Incoming request %s took %v time", r.URL.Path, time.Since(t1)))
	})
}

func sanityCheck() {
	envs := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, e := range envs {
		if os.Getenv(e) == "" {
			log.Fatalf("%s environment varaible missing, terminating application\n", e)
		}
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
