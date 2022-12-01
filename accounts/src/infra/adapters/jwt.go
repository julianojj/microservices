package adapters

import (
	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct{}

func NewJwt() *Jwt {
	return &Jwt{}
}

func (j *Jwt) Encode(data map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": data["Id"],
	})

	tokenString, err := token.SignedString([]byte("1234"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
