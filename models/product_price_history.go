package models

import (
	"encoding/json"
)

type ProductPriceHistory struct {
	Id   string          `json:"id"`
	Data json.RawMessage `json:"data"`
}
