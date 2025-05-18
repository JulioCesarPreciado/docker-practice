package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"example-api/pkg/logger"

	"github.com/golang-jwt/jwt/v5"
)

// CORSMiddleware sets the required headers to allow cross-origin requests.
// It is meant to be used only in development or for specific public endpoints.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware is a middleware that logs request metadata and execution time.
// It delegates log writing to the logger.Info function, which writes to a configured log file.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger.Info(
			fmt.Sprintf("[📡 %s] %s %s (%s)", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start)),
		)

	})
}

// RequireJWT returns a middleware that validates a JWT token using the provided secret.
// If the token is valid, the request is passed to the next handler. Otherwise, it returns a 401 response.
func RequireJWT(secret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secret), nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
