package main

import (
	"fmt"
	"modular-acai-shop/cmd/factory"
	"net/http"
)

func main() {
	app := factory.NewApplication()
	app.LoadEnv()
	defer app.Pg.Close()

	http.HandleFunc("POST /auth/signup", app.UserController.CreateUser)
	http.HandleFunc("POST /auth/signin", app.UserController.SignIn)
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
