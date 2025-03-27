package client

import (
	"time"
	es "api-stock/models/establishment"
)

type RoleClient int

const (
	Admin RoleClient = iota
	Property
	Employee
)

type Client struct {
	Id             string          `json:"id"`
	Email          string          `json:"email"`
	CUIT           string          `json:"cuit"`
	Name           string          `json:"name"`
	Password       string          `json:"password"`
	Cellphone      string          `json:"cellphone"`
	Role           RoleClient      `json:"role"`
	IsActive       bool            `json:"is_active"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	Establishments []es.Establishment `json:"establishments"`
}
