package daily_movement

import (
	"time"
)

type TypeMovement int

const (
	SaleMovement TypeMovement = iota
	BuyMovement
	DiscontinuedMovement
)

type DailyMovement struct {
	Id        string       `json:"id"`
	ProductId string       `json:"product_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Cost      float32      `json:"cost"`
	Price     float32      `json:"price"`
	Movement  TypeMovement `json:"movement"`
}
