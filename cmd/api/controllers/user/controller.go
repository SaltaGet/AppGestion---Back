package user

import (
	"appGestion/pkg/ports"
)

type Controller struct {
	UserService ports.UserService
}