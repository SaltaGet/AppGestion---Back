package utils

import (
	"appGestion/pkg/models/establishment"
	"appGestion/pkg/models/user"
	"os"
	"strings"
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

func VerifyToken(tokenString string) (jwt.Claims, error) {
	cleanToken := CleanToken(tokenString)
	token, err := jwt.Parse(cleanToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}

func GenerateTenantToken(establishment *establishment.Establishment) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": establishment.Id,
	})

	t, err := token.SignedString([]byte(os.Getenv("SECRET_TENANT_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyTenantToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_TENANT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}

func CleanToken(bearerToken string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(bearerToken, prefix) {
		return strings.TrimPrefix(bearerToken, prefix)
	}
	return bearerToken
}
