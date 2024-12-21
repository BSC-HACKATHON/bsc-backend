package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/bsc-auth/config"
)

func GenerateJWT(id uint, email string, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"name":  name,
	})
	tokenString, err := token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
