package user

import (
	"api-stock/pkg/ports"	
)

type Service struct {
	UserRepository ports.UserRepository	
}