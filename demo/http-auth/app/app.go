package app

import (
	"learning-http-auth/domain"
	"learning-http-auth/logger"
	"learning-http-auth/service"
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
	repo := domain.NewAuthRepository(dbConn)
	svc := service.NewLoginService(repo)
	ah := AuthHandler{svc}

	r.HandleFunc("/auth/login", ah.loginHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/verify", ah.verifyHandler).Methods(http.MethodGet)

	logger.Info("starting auth server ....")

	log.Fatal(http.ListenAndServe("localhost:8082", r))
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
