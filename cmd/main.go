package main

import (
	"fmt"
	"modular-acai-shop/internal/auth/application/controller"
	"modular-acai-shop/internal/auth/application/usecase"
	"modular-acai-shop/internal/auth/infra/postgresql/repository"
	"modular-acai-shop/pkg/database"
	"net/http"
)

func main() {
	pg, err := database.NewPostgresDB("postgresql://root:example@localhost:5432/db")
	if err != nil {
		panic(err)
	}
	defer pg.Close()
	conn, err := pg.GetDB()
	if err != nil {
		panic(err)
	}
	repo := repository.NewPostgresUserRepository(conn)
	createUserUseCase := usecase.NewCreateUserUseCase(repo)
	createUserController := controller.NewUserController(createUserUseCase)

	http.HandleFunc("POST /auth/signup", createUserController.CreateUser)
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
