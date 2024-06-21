package factory

import (
	"log"
	"modular-acai-shop/internal/auth/application/controller"
	"modular-acai-shop/internal/auth/application/service"
	"modular-acai-shop/internal/auth/application/usecase"
	"modular-acai-shop/internal/auth/infra/postgresql/repository"
	"modular-acai-shop/pkg/database"
	"net/http"

	"github.com/joho/godotenv"
)

type Application struct {
	Pg             *database.Postgres
	UserController *controller.UserController
}

func ConnectToDb() (*database.Postgres, error) {
	pg, err := database.NewPostgresDB("postgresql://root:example@localhost:5432/db")
	if err != nil {
		panic(err)
	}

	return pg, nil
}

func (a *Application) LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *Application) RunServer(port string) error {
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

	return nil
}

func NewApplication() *Application {
	Pg, err := ConnectToDb()
	if err != nil {
		panic(err)
	}
	conn, err := Pg.GetDB()
	if err != nil {
		panic(err)
	}
	j := service.NewJwtService()
	r := repository.NewPostgresUserRepository(conn)
	cu := usecase.NewCreateUserUseCase(r)
	su := usecase.NewSignInUserUseCase(r, j)
	c := controller.NewUserController(cu, su)
	return &Application{Pg: Pg, UserController: c}
}
