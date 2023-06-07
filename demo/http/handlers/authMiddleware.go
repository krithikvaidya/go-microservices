package handlers

import (
	"learning-http/domain"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

// RBAC -> role based access control

func (am *AuthMiddleware) AuthMiddlewareHandler(next http.Handler) http.Handler {
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
			isAuthorized := am.repo.IsAuthorized(token)
			if isAuthorized {
				next.ServeHTTP(w, r)
			} else {
				writeResponse(w, http.StatusUnauthorized, "Unauthorized")
			}

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

func NewAuthMiddleware(repo domain.AuthRepository) AuthMiddleware {
	return AuthMiddleware{repo}
}
