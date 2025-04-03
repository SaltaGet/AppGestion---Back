package user

import (
	uc "api-stock/pkg/models/user"
)

func (r *Repository) Insert(user *uc.UserCreate) (string, error) {
	return "", nil	
}

func (r *Repository) Update(user *uc.UserUpdate) (error) {
	return nil	
}

func (r *Repository) ChangeStatus(user *uc.UserCreate) (string, error) {
	return "", nil	
}

func (r *Repository) DeleteRelation(user *uc.UserCreate) (string, error) {
	return "", nil	
}

// func (r *Repository) GetByIdentifier(credentials *auth.AuthLogin) (*user.User, error) {
// 	query := `SELECT * FROM users WHERE identifier = ? LIMIT 1`
// 	row, err := database.GetRows(r.DB, query, credentials.Identifier)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var user []user.User
// 	err = utils.MapRowsToStruct(row, &user)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user[0], nil
// }