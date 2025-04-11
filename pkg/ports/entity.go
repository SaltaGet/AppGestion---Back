package ports

import (
	"appGestion/pkg/models/entity"
)

type EntityService interface {
	Create(entity *entity.EntityCreate) (id string, err error)
	GetAll() (entities *[]entity.Entity, err error)
	Update(entity *entity.EntityUpdate) (err error)
}

type EntityRepository interface {
	Insert(entity *entity.EntityCreate) (id string, err error)
	Update(entity *entity.EntityUpdate) (err error)
	ExistById(id string) (exist bool, err error)
	ExistByCUIT(cuit string) (exist bool, err error)
	GetAll() (entities *[]entity.Entity, err error)
}