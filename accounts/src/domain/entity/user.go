package entity

import (
	"errors"
	"net/mail"
)

type User struct {
	Id       string
	Name     string
	Username string
	Email    string
	Password string
}

func NewUser(
	id string,
	name string,
	username string,
	email string,
	password string,
) (*User, error) {
	if isInvalidEmail(email) {
		return nil, errors.New("invalid email")
	}
	if isInvalidPassword(password) {
		return nil, errors.New("invalid password")
	}
	return &User{
		id,
		name,
		username,
		email,
		password,
	}, nil
}

func isInvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func isInvalidPassword(password string) bool {
	return len(password) < 6
}
