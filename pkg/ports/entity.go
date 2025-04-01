package ports

import(
	ent "api-stock/pkg/models/entity"
)

type EntityService interface {
	Create(entity *ent.EntityCreate) (id string, err error)
}

type EntityRepository interface {
	Insert(entity *ent.EntityCreate) (id string, err error)
}