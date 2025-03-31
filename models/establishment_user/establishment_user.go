package establishmentuser

import "time"

type EstablishmentUser struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	EstablishmentId string `json:"establishment_id"`
	RoleId string `json:"role_id"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
