package models

import (
	"time"
)

type UnitMeasure int

const (
	Gramos UnitMeasure = iota
	CentimetrosCubicos
	Unidad
	SinUnidad
)

type Product struct {
	Id            string          `json:"id"`
	SKU           string          `json:"sku"`
	Name          string          `json:"name"`
	UnitMeasure   UnitMeasure     `json:"unit_measure"`
	CategoryId    string          `json:"category_id"`
	Cost          float32         `json:"cost"`
	Price         float32         `json:"price"`
	DateFrom      time.Time       `json:"date_from"`
	DateTo        time.Time       `json:"date_to"`
	Category      Category        `json:"category"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	Sales         []Sale          `json:"sales"`
	Buys          []Sale          `json:"buys"`
	Discontinued  []Sale          `json:"discontinued"`
	MovementToday []DailyMovement `json:"movement_today"`
}
