package entityuser

import (
	"time"
)

type EntityUser struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	EntityId string `json:"entity_id"`
	RoleId string `json:"role_id"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


