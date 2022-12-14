package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateUserIfIdIsRequired(t *testing.T) {
	user, err := NewUser(
		"",
		"any_user",
		"any_username",
		"any_email@test.com",
		"any_password",
	)
	assert.Nil(t, user)
	assert.EqualError(t, err, "id is required")
}

func TestNotShouldCreateUserIfNameIsRequired(t *testing.T) {
	user, err := NewUser(
		uuid.NewString(),
		"",
		"any_username",
		"any_email@test.com",
		"any_password",
	)
	assert.Nil(t, user)
	assert.EqualError(t, err, "name is required")
}

func TestNotShouldCreateUserIfUsernameIsRequired(t *testing.T) {
	user, err := NewUser(
		uuid.NewString(),
		"any_user",
		"",
		"any_email@test.com",
		"any_password",
	)
	assert.Nil(t, user)
	assert.EqualError(t, err, "username is required")
}

func TestNotShouldCreateUserIfEmailIsRequired(t *testing.T) {
	user, err := NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"",
		"any_password",
	)
	assert.Nil(t, user)
	assert.EqualError(t, err, "email is required")
}

func TestNotShouldCreateUserIfPasswordIsRequired(t *testing.T) {
	user, err := NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_email@test.com",
		"",
	)
	assert.Nil(t, user)
	assert.EqualError(t, err, "password is required")
}

func TestNotShouldCreateUserIfInvalidEmail(t *testing.T) {
	user, err := NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_emailtest.com",
		"any_password",
	)
	assert.Nil(t, user)
	assert.EqualError(t, err, "invalid email")
}

func TestNotShouldCreateUserIfInvalidPassword(t *testing.T) {
	user, err := NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_email@test.com",
		"12345",
	)
	assert.Nil(t, user)
	assert.EqualError(t, err, "invalid password")
}

func TestCreateUser(t *testing.T) {
	user, err := NewUser(
		uuid.NewString(),
		"any_user",
		"any_username",
		"any_email@test.com",
		"any_password",
	)
	assert.Equal(t, user.Name, "any_user")
	assert.NoError(t, err)
}
