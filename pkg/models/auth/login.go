package auth

import "github.com/go-playground/validator/v10"

type AuthLogin struct {
	Identifier string `json:"identifier" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (a *AuthLogin) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
