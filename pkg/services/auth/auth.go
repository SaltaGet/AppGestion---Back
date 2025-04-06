package auth

import (
	"api-stock/pkg/models/auth"
	"api-stock/pkg/models"
	"api-stock/pkg/models/user"
	"api-stock/pkg/utils"
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
	user, err := s.UserRepository.GetByIdentifier(userId)

	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener usuario", err)
	}

	if user == nil {
		return nil, models.ErrorResponse(404, "Usuario no encontrado", err)
	}

	return user,nil
}