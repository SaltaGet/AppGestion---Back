package auth

import (
	"api-stock/pkg/models/auth"
	"api-stock/pkg/utils"
)

func (s *Service) Login(credentials *auth.AuthLogin) (string, int, error) {
	user, err := s.UserRepository.GetByIdentifier(credentials.Identifier)

	if err != nil {
		return "Error al intentar conseguir el usuario", 500, err
	}

	if user == nil {
		return "Usuario no encontrado", 404 ,err
	}

	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		return "Credenciales incorrecta", 401, err
	}

	token, err := utils.GenerateUserToken(user)
	if err != nil {
		return "Error al generar el token", 500, err
	}

	return token, 200, nil
}

// import(
// 	"api-stock/database"
// 	u "api-stock/core/models/user"
// )

// func GetCurrentUser(userId string) (*u.User, error){
// 	query := `
// 		SELECT 
// 			u.id, u.first_name, u.last_name, u.email, u.identifier, 
// 			u.phone, u.address, u.city, u.country, u.zip_code, 
// 			u.password, u.created, u.updated, r.id, r.name, r.hierarchy
// 		FROM users u
// 		JOIN entity_user eu ON u.id = eu.user_id
// 		JOIN roles r ON eu.entity_id = r.entity_id
// 		WHERE u.id = ?
// 	`

// 	row :=database.GetRow(query, userId)

// 	var user u.User
// 	err := row.Scan(
// 		&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Identifier,
// 		&user.Phone, &user.Address, &user.City, &user.Country, &user.ZipCode,
// 		&user.Password, &user.Created, &user.Updated,
// 		&user.Role.Id, &user.Role.Name, &user.Role.Hierarchy,
// 	)

// 	if err != nil {
//     return nil, err
// 	}

// 	return &user, nil
// }