package usecase

import (
	"testing"

	"github.com/julianojj/microservices/accounts/src/infra/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateUserIfUserAlreadyExists(t *testing.T) {
	userRepository := memory.NewCreateUserRepository()
	createUser := NewCreateUser(userRepository)
	inputCreateUser := CreateUserInput{
		"any_user",
		"any_username",
		"any_email@test.com",
		"any_password",
	}
	createUser.Execute(inputCreateUser)
	_, err := createUser.Execute(inputCreateUser)
	assert.EqualError(t, err, "user already exists")
}

func TestShouldCreateUser(t *testing.T) {
	userRepository := memory.NewCreateUserRepository()
	createUser := NewCreateUser(userRepository)
	inputCreateUser := CreateUserInput{
		"any_user",
		"any_username",
		"any_email@test.com",
		"any_password",
	}
	outputCreateUser, _ := createUser.Execute(inputCreateUser)
	user := userRepository.Find(outputCreateUser.id)
	assert.Equal(t, user.Name, "any_user")
}
