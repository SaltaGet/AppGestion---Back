package models

import (
	"time"
)

type RoleUser int

const (
	Admin RoleUser = iota
	Property
	Employee
)

type Client struct {
	Id        string    `json:"id"`
	Userame   string    `json:"username"`
	CUIL      string    `json:"cuil"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Cellphone string    `json:"cellphone"`
	Password  string    `json:"password"`
	Role      RoleUser  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
