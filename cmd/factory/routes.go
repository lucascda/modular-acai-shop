package factory

import "net/http"

func RegisterRoutes(app *Application) {
	http.HandleFunc("POST /auth/signup", app.UserController.CreateUser)
	http.HandleFunc("POST /auth/signin", app.UserController.SignIn)
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
}
