package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

type AuthMiddleware struct{}

func (ah *AuthMiddleware) AuthMiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// retrieve the token from the request header
		// verify the token from the http://localhost:8082/auth/verify auth api
		// if success then forward to handler, in case of error return 401

		// sample header
		// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYxMjQ1OTcsInVzZXJJZCI6IjIwMDAifQ.cGQLkbx8YAULQgNtkSUgYEwjhzj5ocCKRhWKcS8jvYc
		token := getTokenFromHeader(r.Header.Get("Authorization"))
		if token == "" {
			writeResponse(w, http.StatusUnauthorized, "missing token")
		} else {

			// make the request to /auth/verify?token=some-jwt-token

			fmt.Println(token)
			next.ServeHTTP(w, r)
		}
	})
}

func getTokenFromHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}
	splitToken := strings.Split(authHeader, "Bearer")
	return strings.TrimSpace(splitToken[1])
}

func NewAuthMiddleware() AuthMiddleware {
	return AuthMiddleware{}
}
