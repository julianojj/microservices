package main

import (
	"fmt"
	"net/http"

	"github.com/julianojj/microservices/accounts/src/application/usecase"
	"github.com/julianojj/microservices/accounts/src/infra/api/controller"
	"github.com/julianojj/microservices/accounts/src/infra/api/route"
	"github.com/julianojj/microservices/accounts/src/infra/repository/memory"
)

func main() {
	app := http.NewServeMux()
	port := 8000
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: app,
	}
	userRepository := memory.NewCreateUserRepository()
	createUser := usecase.NewCreateUser(userRepository)
	createUserController := controller.NewCreateUserController(createUser)
	route.NewUserRoute(
		app, 
		createUserController,
	).Init()
	server.ListenAndServe()
}
