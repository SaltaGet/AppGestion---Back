package establishment

import (
	"time"
)

type Establishment struct {
	Id          string    `json:"id"`
	Name string    `json:"name"`
	Email     string `json:"email"`
	Phone string `json:"phone"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	ZipCode     string    `json:"zip_code"`
	DateCreate   time.Time `json:"date_create"`
	Connection  string    `json:"connection"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	EntityId    string    `json:"entity_id"`
}
