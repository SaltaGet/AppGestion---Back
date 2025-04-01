package entity

import (
	"api-stock/pkg/ports"
)

type Controller struct {
	EntityService ports.EntityService
}