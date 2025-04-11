package auth

import (
	"appGestion/pkg/models/auth"
	"appGestion/pkg/models"
	"appGestion/pkg/models/user"
	"appGestion/pkg/utils"
)

func (s *Service) Login(credentials *auth.AuthLogin) (string, error) {
	user, err := s.UserRepository.GetByIdentifier(credentials.Identifier)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al obtener usuario", err)
	}

	if user == nil {
		return "", models.ErrorResponse(404, "Usuario no encontrado", err)
	}

	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		return "", models.ErrorResponse(401, "Credenciales incorrectas", err)
	}

	token, err := utils.GenerateUserToken(user)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al generar token", err)
	}

	return token, nil
}

func (s *Service) GetCurrentUser(userId string) (*user.User, error){
	user, err := s.UserRepository.GetById(userId)

	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener usuario", err)
	}

	if user == nil {
		return nil, models.ErrorResponse(404, "Usuario no encontrado", err)
	}

	return user,nil
}

func (s *Service) GetConnectionTenant(establishmentId string, userId string) (string, error) {
	if establishmentId == "" || userId == "" {
		return "", models.ErrorResponse(400, "El id del establecimiento o el id del usuario no pueden estar vacíos", nil)
	}

	connection, err := s.EstablishmentRepository.GetEstablishmentById(establishmentId, userId)
	
	if err != nil {
		return "", models.ErrorResponse(500, "Error al obtener establecimiento", err)
	}

	if connection == "" {
		return "", models.ErrorResponse(404, "Establecimiento no encontrado", err)
	}

	uri, err := utils.Decrypt(connection)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al desencriptar la uri", err)
	}
	
	return uri, nil
}