package entity

type EntityUpdate struct {
	Email           string              `json:"email"`
	Name            string              `json:"name"`
	Phone           string              `json:"phone"`
	Address         string              `json:"address"`
	City            string              `json:"city"`
	Country         string              `json:"country"`
	ZipCode         string              `json:"zip_code"`
}