package entity

import (
	"appGestion/pkg/ports"
)

type Controller struct {
	EntityService ports.EntityService
}