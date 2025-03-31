package services

import(
	"api-stock/database"
	u "api-stock/models/user"
)

func GetCurrentUser(userId string) (*u.User, error){
	query := `
		SELECT 
			u.id, u.first_name, u.last_name, u.email, u.identifier, 
			u.phone, u.address, u.city, u.country, u.zip_code, 
			u.password, u.created, u.updated, r.id, r.name, r.hierarchy
		FROM users u
		JOIN entity_user eu ON u.id = eu.user_id
		JOIN roles r ON eu.entity_id = r.entity_id
		WHERE u.id = ?
	`

	row :=database.GetRow(query, userId)

	var user u.User
	err := row.Scan(
		&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Identifier,
		&user.Phone, &user.Address, &user.City, &user.Country, &user.ZipCode,
		&user.Password, &user.Created, &user.Updated,
		&user.Role.Id, &user.Role.Name, &user.Role.Hierarchy,
	)

	if err != nil {
    return nil, err
	}

	return &user, nil
}