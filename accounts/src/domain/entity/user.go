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
	user := &User{
		id,
		name,
		username,
		email,
		password,
	}
	err := user.IsValid()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) IsValid() error {
	if u.Id == "" {
		return errors.New("id is required")
	}
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if isInvalidEmail(u.Email) {
		return errors.New("invalid email")
	}
	if isInvalidPassword(u.Password) {
		return errors.New("invalid password")
	}
	return nil
}

func isInvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func isInvalidPassword(password string) bool {
	return len(password) < 6
}
