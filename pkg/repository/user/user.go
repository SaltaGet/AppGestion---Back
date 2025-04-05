package user

import (
	"api-stock/pkg/models/user"
	"api-stock/pkg/repository/database"
	"api-stock/pkg/utils"
)

// func (r *Repository) Insert(user *user.UserCreate) (string, error) {
// 	return "", nil	
// }

// func (r *Repository) Update(user *user.UserUpdate) (error) {
// 	return nil	
// }

// func (r *Repository) ChangeStatus(user *user.UserCreate) (string, error) {
// 	return "", nil	
// }

// func (r *Repository) DeleteRelation(user *user.UserCreate) (string, error) {
// 	return "", nil	
// }

func (r *Repository) GetByIdentifier(identifier string) (*user.User, error) {
	query := `SELECT * FROM users WHERE identifier = ? LIMIT 1`
	row, err := database.GetRows(r.DB, query, identifier)

	if err != nil {
		return nil, err
	}

	var user []user.User
	err = utils.MapRowsToStruct(row, &user)

	if err != nil {
		return nil, err
	}

	return &user[0], nil
}

func (r *Repository) GetById(id string) (*user.User, error) {
	query := `SELECT * FROM users WHERE id = ? LIMIT 1`
	row, err := database.GetRows(r.DB, query, id)

	if err != nil {
		return nil, err
	}

	var user []user.User
	err = utils.MapRowsToStruct(row, &user)

	if err != nil {
		return nil, err
	}

	return &user[0], nil
}

func (r *Repository) GetByEmail(email string) (*user.User, error) {
	query := `SELECT * FROM users WHERE email = ? LIMIT 1`
	row, err := database.GetRows(r.DB, query, email)

	if err != nil {
		return nil, err
	}

	var user []user.User
	err = utils.MapRowsToStruct(row, &user)

	if err != nil {
		return nil, err
	}

	return &user[0], nil
}