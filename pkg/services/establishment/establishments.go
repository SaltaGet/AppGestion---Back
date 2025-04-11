package establishment

import (
	"appGestion/pkg/models/establishment"
	"appGestion/pkg/models"
	"appGestion/pkg/repository/database"
	"os"
	"fmt"
	"appGestion/pkg/utils"
)

func (s *Service) Create(establishment *establishment.EstablishmentCreate) (string, error) {
	exist, err := s.EntityRepository.ExistById(establishment.EntityId)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al obtener entidad", err)
	}

	if !exist {
		return "", models.ErrorResponse(400, "La entidad no existe", err)
	}

	uri := fmt.Sprintf("%s%s%s%s",os.Getenv("URI_PATH"),establishment.Name,establishment.EntityId,os.Getenv("URI_CONFIG"))
	connection, err := utils.Encrypt(uri)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al encriptar", err)
	}

	err = database.PrepareDB(uri)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al generar base de datos del tenant", err)
	}

	newId, err := s.EstablishmentRepository.Create(establishment, connection)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear establecimiento", err)
	}
	
	return newId, nil
}