package auth

import (
	"api-stock/pkg/models/auth"
	"api-stock/pkg/models/user"
	// "api-stock/pkg/repository/database"
	// "api-stock/pkg/utils"
)

func (r *Repository) Login(credentials *auth.AuthLogin) (*user.User, error) {
	// query := `SELECT * FROM users WHERE identifier = ? LIMIT 1`
	// row, err := database.GetRows(r.DB, query, credentials.Identifier)

	// if err != nil {
	// 	return nil, err
	// }

	// var user user.User
	// err = utils.MapRowToStruct(row, &user)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}