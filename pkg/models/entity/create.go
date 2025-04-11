package entity

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type EntityCreate struct {
	Email           string `json:"email" validate:"required,max=100,email" example:"entity@name.com"`
	CUIT            string `json:"cuit" validate:"required,max=20" example:"20123456789"`
	Name            string `json:"name" validate:"required,max=100" example:"Entity Name"`
	Phone           string `json:"phone" validate:"required,max=20,numeric" example:"1234567890"`
	StartActitivies string `json:"start_activities" validate:"required,dateonly" example:"2023-01-01"`
	Address         string `json:"address" validate:"max=100" example:"123 Main St"`
	City            string `json:"city" validate:"max=100" example:"New York"`
	Country         string `json:"country" validate:"max=100" example:"USA"`
	ZipCode         string `json:"zip_code" validate:"max=20" example:"10001"`
}

func (c *EntityCreate) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("dateonly", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("2006-01-02", fl.Field().String())
		return err == nil
	})

	return validate.Struct(c)
}
