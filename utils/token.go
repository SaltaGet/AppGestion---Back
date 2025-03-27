package utils

import (
	"os"
	cl "api-stock/models/client"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(client *cl.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cuit": client.CUIT,
	})

	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
