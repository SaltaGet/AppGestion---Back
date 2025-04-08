package establishment

import (
	"appGestion/pkg/ports"
)

type Controller struct {
	EstablishmentService ports.EstablishmentService
}