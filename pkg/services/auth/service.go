package auth

import (
	"api-stock/pkg/ports"	
)

type Service struct {
	AuthRepository ports.AuthRepository	
	UserRepository ports.UserRepository
}