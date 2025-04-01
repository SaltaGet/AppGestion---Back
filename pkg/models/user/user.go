package user

import (
	"time"
	ent "api-stock/pkg/models/entity"
	est "api-stock/pkg/models/establishment"
	r "api-stock/pkg/models/role"
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
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Entities []ent.Entity `json:"entities"`
	Establishments []est.Establishment `json:"establishments"`
	Role r.Role `json:"role"`
}
