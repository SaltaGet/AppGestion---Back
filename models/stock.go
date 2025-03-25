package models

import (
)

type Stock struct {
	Id        string    `json:"id"`
	ProductId   string    `json:"product_id"`
	Stock float32    `json:"stock"`
}