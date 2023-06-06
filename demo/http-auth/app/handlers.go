package app

import (
	"encoding/json"
	"learning-http-auth/domain"
	"learning-http-auth/service"
	"learning-http/logger"
	"net/http"
)

type AuthHandler struct {
	service service.DefaultAuthService
}

func (ah *AuthHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req domain.Login
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Info("error while decoding request..." + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, appErr := ah.service.Login(req)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr)
		} else {
			writeResponse(w, http.StatusOK, map[string]string{"token": token})
		}
	}

}

func (ah *AuthHandler) verifyHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
