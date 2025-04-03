package ports

import(
	"api-stock/pkg/models/user"
)

type UserService interface {
	Create(user *user.UserCreate) (id string, err error)
	Update(user *user.UserUpdate) (err error)
}

type UserRepository interface {
	Insert(user *user.UserCreate) (id string, err error)
	Update(user *user.UserUpdate) (err error)
}