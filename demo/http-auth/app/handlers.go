package app

import (
	"encoding/json"
	"learning-http-auth/service"
	"net/http"
)

type AuthHandler struct {
	service service.DefaultAuthService
}

func (ah *AuthHandler) loginHandler(w http.ResponseWriter, r *http.Request) {

}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
