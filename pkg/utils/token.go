package utils

import (
	"os"
	"api-stock/pkg/models/user"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateUserToken(user *user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"first_name": user.FirstName,
		"last_name": user.LastName,
		"email": user.Email,
		"identifier": user.Identifier,
	})

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyClientToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
