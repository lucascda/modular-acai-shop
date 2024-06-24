package factory

import (
	"modular-acai-shop/pkg/middleware"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(middleware.UserContextKey).(string)
	if !ok {
		http.Error(w, "no user id", http.StatusUnauthorized)
	}
	w.Write([]byte(id))
}

func RegisterRoutes(app *Application) {
	http.HandleFunc("POST /auth/signup", app.UserController.CreateUser)
	http.HandleFunc("POST /auth/signin", app.UserController.SignIn)
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	http.HandleFunc("GET /me", middleware.AuthMiddleware(http.HandlerFunc(Me)))
}
