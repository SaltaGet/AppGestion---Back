package entity

import (
	// es "api-stock/models/establishment"
	"time"
)

type RoleClient int

const (
	Admin RoleClient = iota
	Property
	Employee
)

type Entity struct {
	Id              string             `json:"id"`
	Email           string             `json:"email"`
	CUIT            string             `json:"cuit"`
	Name            string             `json:"name"`
	Phone           string             `json:"phone"`
	StartActitivies time.Time          `json:"start_activities"`
	Address         string             `json:"address"`
	City            string             `json:"city"`
	Country         string             `json:"country"`
	ZipCode         string             `json:"zip_code"`
	IsActive        bool               `json:"is_active"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	// Establishments  []es.Establishment `json:"establishments"`
}
