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

// http://authserver:8082/auth/verify?token=jwt-token
/*
	{ // good response
		"isAuthorized": true,
	}
	{ // bad response
		"isAuthorized": false,
		"msg": "reason here"
	}
*/
func (ah *AuthHandler) verifyHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	if token != "" {
		appErr := ah.service.Verify(token)
		if appErr != nil {
			writeResponse(w, appErr.Code, notAuthorized(appErr.Message))
		} else {
			// good case
			writeResponse(w, http.StatusOK, map[string]any{"isAuthorized": true})
		}
	} else {
		writeResponse(w, http.StatusForbidden, notAuthorized("token missing"))
	}
}

func notAuthorized(msg string) map[string]any {
	return map[string]any{"isAuthorized": false, "msg": msg}
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
