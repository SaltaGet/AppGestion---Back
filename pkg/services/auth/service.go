package auth

import (
	"appGestion/pkg/ports"	
)

type Service struct {
	AuthRepository ports.AuthRepository	
	UserRepository ports.UserRepository
	EstablishmentRepository ports.EstablishmentRepository
}
