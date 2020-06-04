package middlewares

import (

	"log"
	"strings"
	"encoding/json"

	"net/http"
)

// SimpleMiddleware - Middleware for just pass
func SimpleMiddleware(next http.Handler) http.Handler {

	log.Println("Simple Middleware!!!")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.RequestURI)

		next.ServeHTTP(w, r)

	})
}

// AuthMiddleware - validating email and token value
func AuthMiddleware(next http.Handler) http.Handler {

	log.Println("Auth Middleware!!!")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("validate-token")
		email := r.Header.Get("auth_email")
		token = strings.TrimSpace(token)

		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}

		if email == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
