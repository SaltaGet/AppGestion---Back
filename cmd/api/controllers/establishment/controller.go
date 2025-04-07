package establishment

import (
	"api-stock/pkg/ports"
)

type Controller struct {
	EstablishmentService ports.EstablishmentService
}