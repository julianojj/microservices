package usecase

import (
	"errors"

	"github.com/julianojj/microservices/accounts/internal/domain/repository"
	"github.com/julianojj/microservices/accounts/internal/infra/adapters"
)

type AuthenticateUser struct {
	UserRepository repository.UserRepository
	Hash           adapters.Hash
	Sign           adapters.Sign
}

type AuthenticateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUserOutput struct {
	Token string `json:"token"`
}

func NewAuthenticateUser(
	userRepository repository.UserRepository,
	hash adapters.Hash,
	sign adapters.Sign,
) *AuthenticateUser {
	return &AuthenticateUser{
		UserRepository: userRepository,
		Hash:           hash,
		Sign:           sign,
	}
}

func (a *AuthenticateUser) Execute(input AuthenticateUserInput) (*AuthenticateUserOutput, error) {
	if input.Email == "" {
		return nil, errors.New("email is required")
	}
	if input.Password == "" {
		return nil, errors.New("password is required")
	}
	existingUser := a.UserRepository.FindByEmail(input.Email)
	if existingUser == nil {
		return nil, errors.New("invalid login")
	}
	err := a.Hash.Decrypt([]byte(existingUser.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid login")
	}
	token, err := a.Sign.Encode(map[string]interface{}{
		"Id": existingUser.Id,
	})
	if err != nil {
		return nil, err
	}
	return &AuthenticateUserOutput{
		Token: token,
	}, nil
}
