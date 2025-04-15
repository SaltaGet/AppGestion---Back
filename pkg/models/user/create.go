package user

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"unicode"
)

type UserCreate struct {
	FirstName  string `json:"first_name" validate:"required,max=100" example:"John"`
	LastName   string `json:"last_name" validate:"required,max=100" example:"Doe"`
	Email      string `json:"email" validate:"required,max=100,email" example:"john@doe.com"`
	Identifier string `json:"identifier" validate:"required,max=20" example:"12345678"`
	Phone      string `json:"phone" validate:"required,max=20,numeric" example:"1234567890"`
	Address    string `json:"address" validate:"max=100" example:"123 Main St"`
	City       string `json:"city" validate:"max=100" example:"New York"`
	Country    string `json:"country" validate:"max=100" example:"USA"`
	ZipCode    string `json:"zip_code" validate:"max=20" example:"10001"`
	Password   string `json:"password" validate:"required,strongpassword" example:"P@ssw0rd"`
	EntityId string `json:"entity_id" validate:"required" example:"00000000-aaaa-0000-aaaa-000000000000"`
	RoleId string `json:"role_id" validate:"required" example:"00000000-aaaa-0000-aaaa-000000000000"`
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


func (c *UserCreate) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("strongpassword", strongPasswordValidator)
	return validate.Struct(c)
}
