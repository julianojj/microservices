package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user, _ := NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_email@test.com",
		"any_password",
	)
	assert.Equal(t, user.Name, "any_user")
}

func TestNotShouldCreateUserIfInvalidEmail(t *testing.T) {
	_, err := NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_emailtest.com",
		"any_password",
	)
	assert.EqualError(t, err, "invalid email")
}

func TestNotShouldCreateUserIfInvalidPassword(t *testing.T) {
	_, err := NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_email@test.com",
		"12345",
	)
	assert.EqualError(t, err, "invalid password")
}
