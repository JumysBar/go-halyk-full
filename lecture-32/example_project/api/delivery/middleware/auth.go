package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}

		if user != "admin" || password != "password" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Such user not found"))
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
