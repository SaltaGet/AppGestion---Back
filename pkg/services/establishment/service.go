package establishment

import (
	"appGestion/pkg/ports"	
)

type Service struct {
	EntityRepository ports.EntityRepository	
	EstablishmentRepository ports.EstablishmentRepository
}
