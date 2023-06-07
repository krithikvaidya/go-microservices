package app

import (
	"fmt"
	"learning-http-auth/domain"
	"learning-http-auth/logger"
	"learning-http-auth/service"
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
	repo := domain.NewAuthRepository(dbConn)
	svc := service.NewLoginService(repo)
	ah := AuthHandler{svc}

	r.HandleFunc("/auth/login", ah.loginHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/verify", ah.verifyHandler).Methods(http.MethodGet)

	logger.Info("starting auth server ....")

	log.Fatal(http.ListenAndServe("localhost:8082", r))
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
