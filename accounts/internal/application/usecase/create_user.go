package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/julianojj/microservices/accounts/internal/domain/entity"
	"github.com/julianojj/microservices/accounts/internal/domain/repository"
	"github.com/julianojj/microservices/accounts/internal/infra/adapters"
)

type CreateUser struct {
	userRepository repository.UserRepository
	hash           adapters.Hash
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	id string
}

func NewCreateUser(
	userRepository repository.UserRepository,
	hash adapters.Hash,
) *CreateUser {
	return &CreateUser{
		userRepository,
		hash,
	}
}

func (c *CreateUser) Execute(input CreateUserInput) (CreateUserOutput, error) {
	userId := uuid.NewString()
	user, err := entity.NewUser(
		userId,
		input.Name,
		input.Username,
		input.Email,
		input.Password,
	)
	if err != nil {
		return CreateUserOutput{}, err
	}
	existingUser := c.userRepository.FindByEmail(user.Email)
	if existingUser != nil {
		return CreateUserOutput{}, errors.New("user already exists")
	}
	encryptedPassword, _ := c.hash.Encrypt([]byte(user.Password))
	user.Password = string(encryptedPassword)
	c.userRepository.Save(user)
	return CreateUserOutput{
		id: userId,
	}, nil
}
