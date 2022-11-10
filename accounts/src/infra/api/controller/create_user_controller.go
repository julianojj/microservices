package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julianojj/microservices/accounts/src/application/usecase"
)

type CreateUserController struct {
	createUser *usecase.CreateUser
}

func NewCreateUserController(
	createUser *usecase.CreateUser,
) *CreateUserController {
	return &CreateUserController{
		createUser,
	}
}

func (u *CreateUserController) Handle(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInput
	json.NewDecoder(r.Body).Decode(&input)
	_, err := u.createUser.Execute(input)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		return
	}
	if err.Error() == "invalid email" || err.Error() == "invalid password" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err.Error() == "user already exists" {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	http.Error(w, "internal server error", http.StatusInternalServerError)
}
