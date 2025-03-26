package models

import (
	"encoding/json"
	"time"
)

type Discontinued struct {
	Id   string          `json:"id"`
	Date time.Time       `json:"date"`
	Data json.RawMessage `json:"data"`
}
