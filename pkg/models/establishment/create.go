package establishment

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type EstablishmentCreate struct {
	Name       string    `json:"name" validate:"required,max=100"`
	Email      string    `json:"email" validate:"required,email"`
	Phone      string    `json:"phone" valdiate:"required,max=20,numeric"`
	Address    string    `json:"address" validate:"required,max=255"`
	City       string    `json:"city" validate:"required,max=100"`
	Country    string    `json:"country" validate:"required,max=255"`
	ZipCode    string    `json:"zip_code" validate:"required,max=20"`
	DateCreate time.Time `json:"date_create" validate:"required"`
	EntityId   string    `json:"entity_id"`
}

func (e *EstablishmentCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}
