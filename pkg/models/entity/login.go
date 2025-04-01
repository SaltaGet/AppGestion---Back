package entity

import (
	"github.com/go-playground/validator/v10"
)

type ClientLogin struct {
	CUIT      string `json:"cuit" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (c *ClientLogin) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("strongpassword", strongPasswordValidator)
	return validate.Struct(c)
}