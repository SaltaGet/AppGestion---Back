package entity

import (
	// db "api-stock/pkg/repository/database"
	"appGestion/pkg/models/entity"
	"appGestion/pkg/models"
	// "api-stock/pkg/utils"
	
	// "database/sql"
	// "os"
	// "time"
	// "errors"
)


func (s *Service) Create(entity *entity.EntityCreate) (string, error) {
	exist, err := s.EntityRepository.ExistByCUIT(entity.CUIT)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al intentar acceder a la entidad", err)
	}

	if exist {
		return "", models.ErrorResponse(404, "La entidad ya existe", err)
	}

	newId, err := s.EntityRepository.Insert(entity)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al intentar crear la entidad", err)
	}

	return newId, nil
}

func (s *Service) GetAll() (*[]entity.Entity, error) {
	entities, err := s.EntityRepository.GetAll()

	if err != nil {
		return nil, models.ErrorResponse(500, "Error al recuperar las entidades", err)
	}

	return entities, nil
}

func (s Service) Update(entity *entity.EntityUpdate) error {
	return nil
}

func GetClientByCUIT(id string) (bool, error) {
	// var exist bool

	// db := db.GetDB()

	// query := `SELECT EXISTS(SELECT 1 FROM clients WHERE cuit = ?)`

	// err := db.QueryRow(query, id).Scan(&exist)

	// if err != nil {
	// 	return false, err
	// }

	// return exist, nil
	return true, nil
}

func (s *Service) Insert(entity *entity.EntityCreate) (string, error) {
	return "", nil
}
