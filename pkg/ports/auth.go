package ports

import (
	"appGestion/pkg/models/auth"
	"appGestion/pkg/models/user"
)

type AuthService interface {
	Login(credentials *auth.AuthLogin) (token string, err error)
	GetCurrentUser(userId string) (user *user.User, err error)
	GetConnectionTenant(establishmentId string, userId string) (uri string, err error)
}

type AuthRepository interface {
	Login(credentials *auth.AuthLogin) (user *user.User, err error)
}