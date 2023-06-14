package helpers

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/go-oauth2/oauth2/v4/manage"
)

type Middleware struct {
	Manager *manage.Manager
}

type Adapter func(http.Handler) http.Handler

type ContextKey string

const (
	UserIDContextKey ContextKey = "userID"
)

func (m Middleware) Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func (m Middleware) Cors(methods ...string) Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", os.Getenv("CLIENT_URL"))
			w.Header().Add("Vary", "Origin")

			if r.Method == "OPTIONS" {
				w.Header().Add("Access-Control-Allow-Methods", strings.Join(methods, ", "))
				w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (m Middleware) ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/vnd.api+json")

		next.ServeHTTP(w, r)
	})
}

func (m Middleware) Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := header[7:]

		ti, err := m.Manager.LoadAccessToken(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDContextKey, ti.GetUserID())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
