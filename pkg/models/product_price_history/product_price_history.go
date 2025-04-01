package product_price_history

import (
	"encoding/json"
)

type ProductPriceHistory struct {
	Id   string          `json:"id"`
	Data json.RawMessage `json:"data"`
}
