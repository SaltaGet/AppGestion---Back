package product

import (
	buy "api-stock/models/buy"
	cat "api-stock/models/category"
	dm "api-stock/models/daily_movement"
	disc "api-stock/models/discontinued"
	sal "api-stock/models/sale"
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
	Id            string              `json:"id"`
	SKU           string              `json:"sku"`
	Name          string              `json:"name"`
	UnitMeasure   UnitMeasure         `json:"unit_measure"`
	CategoryId    string              `json:"category_id"`
	Cost          float32             `json:"cost"`
	Price         float32             `json:"price"`
	DateFrom      time.Time           `json:"date_from"`
	DateTo        time.Time           `json:"date_to"`
	Category      cat.Category        `json:"category"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	Sales         []sal.Sale          `json:"sales"`
	Buys          []buy.Buy           `json:"buys"`
	Discontinued  []disc.Discontinued `json:"discontinued"`
	MovementToday []dm.DailyMovement  `json:"movement_today"`
}
