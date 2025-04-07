package establishment

import (
	"api-stock/pkg/models/establishment"
	"api-stock/pkg/models"
	"api-stock/pkg/repository/database"
	"os"
	"fmt"
	"api-stock/pkg/utils"
)

func (s *Service) Create(establishment *establishment.EstablishmentCreate) (string, error) {
	exist, err := s.EntityRepository.Exist(establishment.EntityId)

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