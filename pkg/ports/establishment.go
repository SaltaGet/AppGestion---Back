package ports

import(
	"api-stock/pkg/models/establishment"
)

type EstablishmentService interface {
	Create(establishment *establishment.EstablishmentCreate) (id string, err error)
}

type EstablishmentRepository interface {
	Create(establishment *establishment.EstablishmentCreate, connection string) (id string, err error)
}
