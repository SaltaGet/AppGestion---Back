package establishment

import (
	"appGestion/pkg/models"
	"appGestion/pkg/models/establishment"
	"appGestion/pkg/repository/database"
	"appGestion/pkg/utils"
	"fmt"
	"os"
	"strings"
)

func (s *Service) Create(establishment *establishment.EstablishmentCreate) (string, error) {
	exist, err := s.EntityRepository.ExistById(establishment.EntityId)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al obtener entidad", err)
	}

	if !exist {
		return "", models.ErrorResponse(400, "La entidad no existe", err)
	}

	establishmentName := strings.ReplaceAll(establishment.Name, " ", "_")
	uri := fmt.Sprintf("%s%s_%s.db%s",os.Getenv("URI_PATH"),establishmentName,establishment.EntityId,os.Getenv("URI_CONFIG"))
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

func (s *Service) GetAllAdmin() (*[]establishment.Establishment, error) {
	establishments, err := s.EstablishmentRepository.GetAllAdmin()

	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener establecimientos", err)
	}

	return establishments, nil
}