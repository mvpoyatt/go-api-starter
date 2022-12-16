package api

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var key = []byte("my_secret_key")

type Claims struct {
	Email  string
	Expiry time.Time
}

func NewToken(email string) (string, error) {
	expiryTime := time.Now().Add(5 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"expiry": jwt.NewNumericDate(expiryTime),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}