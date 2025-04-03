package auth

import (
	"api-stock/pkg/ports"
)

type Controller struct {
	AuthService ports.AuthService
}