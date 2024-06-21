package controller

import (
	"encoding/json"
	"modular-acai-shop/internal/auth/application/dto"
	"modular-acai-shop/internal/auth/application/usecase"
	"net/http"
)

type Response struct {
	Data any `json:"data"`
}

type UserController struct {
	createUserUseCase *usecase.CreateUserUseCase
	signInUserUseCase *usecase.SignInUserUseCase
}

func NewUserController(createUserUseCase *usecase.CreateUserUseCase, signInUserUseCase *usecase.SignInUserUseCase) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
		signInUserUseCase: signInUserUseCase,
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

func (c *UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data dto.SignInUser
	json.NewDecoder(r.Body).Decode(&data)
	t, err := c.signInUserUseCase.Execute(r.Context(), data.Email, data.Password)
	if err != nil {
		if err.Error() == "unauthorized" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&Response{Data: t})
}
