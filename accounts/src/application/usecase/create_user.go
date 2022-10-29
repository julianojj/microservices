package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/julianojj/microservices/accounts/src/domain/entity"
	"github.com/julianojj/microservices/accounts/src/domain/repository"
)

type CreateUser struct {
	userRepository repository.UserRepository
}

type CreateUserInput struct {
	name     string
	username string
	email    string
	password string
}

type CreateUserOutput struct {
	id string
}

func NewCreateUser(
	userRepository repository.UserRepository,
) *CreateUser {
	return &CreateUser{
		userRepository,
	}
}

func (c *CreateUser) Execute(input CreateUserInput) (CreateUserOutput, error) {
	existingUser := c.userRepository.FindByEmail(input.email)
	if existingUser != nil {
		return CreateUserOutput{}, errors.New("user already exists")
	}
	userId := uuid.NewString()
	user, errUser := entity.NewUser(
		userId,
		input.name,
		input.username,
		input.email,
		input.password,
	)
	if errUser != nil {
		return CreateUserOutput{}, errUser
	}
	c.userRepository.Save(user)
	return CreateUserOutput{
		id: userId,
	}, nil
}
