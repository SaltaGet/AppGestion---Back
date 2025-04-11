package user

import (
	"appGestion/pkg/ports"	
)

type Service struct {
	UserRepository ports.UserRepository	
	EntityRepository ports.EntityRepository
	RoleRepository ports.RoleRepository
}
