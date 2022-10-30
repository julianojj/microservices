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
	json.NewEncoder(w).Encode(input)
}
