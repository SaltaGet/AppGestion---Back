package user

import (
	"appGestion/pkg/models/user"
	"appGestion/pkg/models"
)

func (s *Service) Create(user *user.UserCreate) (string, error) {
	existEntity, err := s.EntityRepository.ExistById(user.EntityId)

	if err != nil {
		return "", models.ErrorResponse(500, "Ha ocurrido un error al validar la entidad", err)
	}

	if !existEntity {
		return "", models.ErrorResponse(404, "Entidad no encontrada", err)
	}

	existRole, err := s.RoleRepository.ExistById(user.RoleId)

	if err != nil {
		return "", models.ErrorResponse(500, "Ha ocurrido un error al validar el rol", err)
	}

	if !existRole {
		return "", models.ErrorResponse(404, "Rol no encontrado", err)
	}

	existUser, err := s.UserRepository.ExistUser(user.Identifier, user.Email)

	if err != nil {
		return "", models.ErrorResponse(500, "Ha ocurrido un error al validar el usuariou", err)
	}

	if existUser {
		return "", models.ErrorResponse(400, "El usuario ya existe", err)
	}

	id, err := s.UserRepository.Create(user)

	if err != nil {
		return "", models.ErrorResponse(500, "Ha ocurrido un error al crear el usuario", err)
	}

	return id, nil
}

func (s *Service) Update(user *user.UserUpdate) error {
	return nil
}