package middleware

import (
	"context"
	"encoding/json"
	"modular-acai-shop/internal/auth/application/service"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey = contextKey("user_id")

type ErrorApiResponse struct {
	Error string `json:"error"`
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h := r.Header.Get("Authorization")
		h_fields := strings.Split(h, " ")
		if h_fields[0] != "Bearer" {
			json.NewEncoder(w).Encode(&ErrorApiResponse{Error: "Token string must start with Bearer"})
			return
		}
		claims, valid := service.NewJwtService().Parse(h_fields[1])

		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&ErrorApiResponse{Error: "Invalid token"})
			return
		}
		sub, err := claims.GetSubject()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&ErrorApiResponse{Error: "Missing subject"})
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, sub)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
