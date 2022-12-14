package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/julianojj/microservices/accounts/internal/domain/entity"
	"github.com/julianojj/microservices/accounts/internal/infra/adapters"
	"github.com/julianojj/microservices/accounts/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldAuthenticateUserIfEmailIsRequired(t *testing.T) {
	userRepository := memory.NewUserRepository()
	hash := adapters.NewBcrypt()
	sign := adapters.NewJwt()
	authenticateUser := NewAuthenticateUser(userRepository, hash, sign)
	input := AuthenticateUserInput{
		"",
		"123456",
	}
	output, err := authenticateUser.Execute(input)
	assert.Nil(t, output)
	assert.EqualError(t, err, "email is required")
}

func TestNotShouldAuthenticateUserIfPasswordIsRequired(t *testing.T) {
	userRepository := memory.NewUserRepository()
	hash := adapters.NewBcrypt()
	sign := adapters.NewJwt()
	authenticateUser := NewAuthenticateUser(userRepository, hash, sign)
	input := AuthenticateUserInput{
		"any_user@test.com",
		"",
	}
	output, err := authenticateUser.Execute(input)
	assert.Nil(t, output)
	assert.EqualError(t, err, "password is required")
}

func TestNotShouldAuthenticateUserIfInvalidEmail(t *testing.T) {
	userRepository := memory.NewUserRepository()
	hash := adapters.NewBcrypt()
	sign := adapters.NewJwt()
	authenticateUser := NewAuthenticateUser(userRepository, hash, sign)
	input := AuthenticateUserInput{
		"any_user@test.com",
		"123456",
	}
	output, err := authenticateUser.Execute(input)
	assert.Nil(t, output)
	assert.EqualError(t, err, "invalid login")
}

func TestNotShouldAuthenticateUserIfInvalidPassword(t *testing.T) {
	userRepository := memory.NewUserRepository()
	hash := adapters.NewBcrypt()
	sign := adapters.NewJwt()
	encryptedPassword, _ := hash.Encrypt([]byte("123456"))
	user, _ := entity.NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_user@test.com",
		"123456",
	)
	user.Password = string(encryptedPassword)
	userRepository.Save(user)
	authenticateUser := NewAuthenticateUser(userRepository, hash, sign)
	input := AuthenticateUserInput{
		"any_user@test.com",
		"1234567",
	}
	output, err := authenticateUser.Execute(input)
	assert.Nil(t, output)
	assert.EqualError(t, err, "invalid login")
}

func TestShouldAuthenticateUser(t *testing.T) {
	userRepository := memory.NewUserRepository()
	hash := adapters.NewBcrypt()
	sign := adapters.NewJwt()
	encryptedPassword, _ := hash.Encrypt([]byte("123456"))
	user, _ := entity.NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_user@test.com",
		"123456",
	)
	user.Password = string(encryptedPassword)
	userRepository.Save(user)
	authenticateUser := NewAuthenticateUser(userRepository, hash, sign)
	input := AuthenticateUserInput{
		"any_user@test.com",
		"123456",
	}
	output, err := authenticateUser.Execute(input)
	assert.NotNil(t, output.Token)
	assert.Nil(t, err)
}
