package adapters

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

func (b *Bcrypt) Encrypt(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func (b *Bcrypt) Decrypt(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
