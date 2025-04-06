package ports

import (
	"api-stock/pkg/models/auth"
	"api-stock/pkg/models/user"
)

type AuthService interface {
	Login(credentials *auth.AuthLogin) (token string, err error)
}

type AuthRepository interface {
	Login(credentials *auth.AuthLogin) (user *user.User, err error)
}