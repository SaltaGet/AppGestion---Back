package client

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"unicode"
)

type ClientCreate struct {
	Email     string `json:"email" validate:"required,email"`
	CUIT      string `json:"cuit" validate:"required,min=10,max=20"`
	Name      string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,strongpassword"`
	Cellphone string `json:"cellphone" validate:"required,max=20,numeric"`
}

func strongPasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	var (
		hasUpper  bool
		hasLower  bool
		hasNumber bool
		hasSymbol bool
	)

	specialCharRegex := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case specialCharRegex.MatchString(string(char)):
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func (c *ClientCreate) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("strongpassword", strongPasswordValidator)
	return validate.Struct(c)
}
