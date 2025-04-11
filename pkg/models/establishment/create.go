package establishment

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type EstablishmentCreate struct {
	Name       string `json:"name" validate:"required,max=100" example:"Establishment Name"`
	Email      string `json:"email" validate:"required,email" example:"name@establishment.com"`
	Phone      string `json:"phone" valdiate:"required,max=20,numeric" example:"1231231231"`
	Address    string `json:"address" validate:"max=100" example:"123 Main St"`
	City       string `json:"city" validate:"max=100" example:"New York"`
	Country    string `json:"country" validate:"max=100" example:"USA"`
	ZipCode    string `json:"zip_code" validate:"max=20" example:"10001"`
	DateCreate string `json:"date_create" validate:"required,dateonly" example:"2023-01-01"`
	EntityId   string `json:"entity_id" example:"00000000-aaaa-0000-aaaa-000000000000"`
}

func (e *EstablishmentCreate) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("dateonly", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("2006-01-02", fl.Field().String())
		return err == nil
	})

	return validate.Struct(e)
}
