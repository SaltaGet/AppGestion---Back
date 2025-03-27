package buy

import (
	"encoding/json"
	"time"
)

type Buy struct {
	Id   string          `json:"id"`
	Date time.Time       `json:"date"`
	Data json.RawMessage `json:"data"`
}
