package ports

import(
	"appGestion/pkg/models/role"
)

type RoleService interface {
	// Create(user *role.Create) (id string, err error)
	// Update(user *user.UserUpdate) (err error)
}

type RoleRepository interface {
	// Insert(user *user.UserCreate) (id string, err error)
	// Update(user *user.UserUpdate) (err error)
	GetById(id string) (role *role.Role, err error) 
	ExistById(id string) (exist bool, err error) 
	GetAll() (roles *[]role.Role, err error) 
}