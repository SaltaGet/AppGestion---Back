package role

import (
	"time"
)

type Role struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Hierarchy int16 `json:"hierarchy"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}