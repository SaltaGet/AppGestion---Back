package establishment

import (
	"api-stock/pkg/ports"	
)

type Service struct {
	EntityRepository ports.EntityRepository	
	EstablishmentRepository ports.EstablishmentRepository
}
