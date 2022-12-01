package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julianojj/microservices/accounts/src/application/usecase"
	"github.com/julianojj/microservices/accounts/src/infra/adapters"
	"github.com/julianojj/microservices/accounts/src/infra/api/controller"
	"github.com/julianojj/microservices/accounts/src/infra/api/route"
	"github.com/julianojj/microservices/accounts/src/infra/repository/memory"
)

func main() {
	mux := http.NewServeMux()
	port := 8000
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	userRepository := memory.NewUserRepository()
	bcrypt := adapters.NewBcrypt()
	createUser := usecase.NewCreateUser(userRepository, bcrypt)
	createUserController := controller.NewCreateUserController(createUser)
	route.NewUserRoute(
		mux,
		createUserController,
	).Init()
	log.Printf("Starting server in %d \n", port)
	server.ListenAndServe()
}
