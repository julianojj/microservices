package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julianojj/microservices/accounts/internal/application/usecase"
	"github.com/julianojj/microservices/accounts/internal/infra/adapters"
	"github.com/julianojj/microservices/accounts/internal/infra/api/controller"
	"github.com/julianojj/microservices/accounts/internal/infra/api/route"
	"github.com/julianojj/microservices/accounts/internal/infra/repository/memory"
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
