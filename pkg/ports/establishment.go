package ports

import(
	"appGestion/pkg/models/establishment"
)

type EstablishmentService interface {
	Create(establishment *establishment.EstablishmentCreate) (id string, err error)
	GetAllAdmin() (establishments *[]establishment.Establishment, err error)
}

type EstablishmentRepository interface {
	Create(establishment *establishment.EstablishmentCreate, connection string) (id string, err error)
	GetEstablishmentById(establishmentId string, userId string) (connection string, err error)
	GetAllAdmin() (establishments *[]establishment.Establishment, err error)
}
