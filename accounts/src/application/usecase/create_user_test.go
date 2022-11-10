package usecase

import (
	"testing"

	"github.com/julianojj/microservices/accounts/src/infra/adapters"
	"github.com/julianojj/microservices/accounts/src/infra/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateUserIfInvalidEmail(t *testing.T) {
	userRepository := memory.NewCreateUserRepository()
	bcrypt := adapters.NewBcrypt()
	createUser := NewCreateUser(userRepository, bcrypt)
	inputCreateUser := CreateUserInput{
		"any_user",
		"any_username",
		"any_emailtest.com",
		"any_password",
	}
	_, err := createUser.Execute(inputCreateUser)
	assert.EqualError(t, err, "invalid email")
}

func TestNotShouldCreateUserIfInvalidPassword(t *testing.T) {
	userRepository := memory.NewCreateUserRepository()
	bcrypt := adapters.NewBcrypt()
	createUser := NewCreateUser(userRepository, bcrypt)
	inputCreateUser := CreateUserInput{
		"any_user",
		"any_username",
		"any_email@test.com",
		"12345",
	}
	_, err := createUser.Execute(inputCreateUser)
	assert.EqualError(t, err, "invalid password")
}

func TestNotShouldCreateUserIfUserAlreadyExists(t *testing.T) {
	userRepository := memory.NewCreateUserRepository()
	bcrypt := adapters.NewBcrypt()
	createUser := NewCreateUser(userRepository, bcrypt)
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
	bcrypt := adapters.NewBcrypt()
	createUser := NewCreateUser(userRepository, bcrypt)
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
