package utils

import (
	"os"
	cl "api-stock/models/client"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateClientToken(client *cl.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cuit": client.CUIT,
	})

	t, err := token.SignedString([]byte(os.Getenv("SECRET_CLIENT_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateUserToken(client *cl.Client) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cuit": client.CUIT,
	})

	t, err := token.SignedString([]byte(os.Getenv("SECRET_USER_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}


func VerifyClientToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_CLIENT_KEY")), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func VerifyUserToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_USER_KEY")), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
