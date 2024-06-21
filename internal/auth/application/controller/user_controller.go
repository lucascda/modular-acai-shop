package controller

import (
	"encoding/json"
	"modular-acai-shop/internal/auth/application/dto"
	"modular-acai-shop/internal/auth/application/usecase"
	"net/http"
)

type UserController struct {
	createUserUseCase *usecase.CreateUserUseCase
}

func NewUserController(createUserUseCase *usecase.CreateUserUseCase) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
	}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var data dto.CreateUser
	json.NewDecoder(r.Body).Decode(&data)
	err := c.createUserUseCase.Execute(r.Context(), data.Name, data.Email, data.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)

}
