package auth

import (
	"appGestion/pkg/ports"
)

type Controller struct {
	AuthService ports.AuthService
}