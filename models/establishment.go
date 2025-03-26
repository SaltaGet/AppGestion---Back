package models

import (
	"time"
)

type Establishment struct {
	Id          string    `json:"id"`
	CompanyName string    `json:"company_name"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	ZipCode     string    `json:"zip_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ClientId    string    `json:"client_id"`
	Connection  string    `json:"connection"`
}
