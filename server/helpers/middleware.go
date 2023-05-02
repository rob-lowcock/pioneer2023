package helpers

import (
	"net/http"
	"os"
	"strings"
)

type Middleware struct {
}

func (m Middleware) Cors(next http.Handler, methods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", os.Getenv("CLIENT_URL"))
		w.Header().Add("Vary", "Origin")

		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Methods", strings.Join(methods, ", "))
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m Middleware) ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/vnd.api+json")

		next.ServeHTTP(w, r)
	})
}
