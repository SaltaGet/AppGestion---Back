package user

import (
	"time"
)

type User struct {
	Id         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Identifier string    `json:"identifier"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	ZipCode    string    `json:"zip_code"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
