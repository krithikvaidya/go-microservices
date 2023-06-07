package domain

import (
	"encoding/json"
	"learning-http/logger"
	"net/http"
	"net/url"
	"os"
)

type AuthRepository struct {
}

func (ar *AuthRepository) IsAuthorized(token string) bool {

	// http://localhost:8082/auth/verify?token=some-token

	authServerHost := os.Getenv("AUTH_SERVER")
	u := url.URL{Host: authServerHost, Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	u.RawQuery = q.Encode()

	// fmt.Sprintf("http://localhost:8080/auth/verify?token=%s", token)

	if response, err := http.Get(u.String()); err != nil {
		logger.Info("Not able to make the auth verify call: " + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err := json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server: " + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

func NewAuthRepository() AuthRepository {
	return AuthRepository{}
}
