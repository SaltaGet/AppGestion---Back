package user

import (
	"api-stock/pkg/ports"
)

type Controller struct {
	UserService ports.UserService
}