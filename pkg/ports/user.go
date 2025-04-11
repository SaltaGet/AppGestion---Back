package ports

import(
	"appGestion/pkg/models/user"
)

type UserService interface {
	Create(user *user.UserCreate) (id string, err error)
	Update(user *user.UserUpdate) (err error)
}

type UserRepository interface {
	Create(user *user.UserCreate) (id string, err error)
	// Update(user *user.UserUpdate) (err error)
	GetByIdentifier(identifier string) (user *user.User, err error) 
	GetById(id string) (user *user.User, err error) 
	GetByEmail(email string) (user *user.User, err error) 
	ExistUser(identifier string, email string) (exist bool, err error) 
}