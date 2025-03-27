package sale

import (
	"encoding/json"
	"time"
)

type Sale struct {
	Id   string          `json:"id"`
	Date time.Time       `json:"date"`
	Data json.RawMessage `json:"data"`
}
