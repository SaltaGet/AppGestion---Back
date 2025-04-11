package user

import (
	"appGestion/pkg/models/user"
	"appGestion/pkg/repository/database"
	"appGestion/pkg/utils"

	"github.com/google/uuid"
)

func (r *Repository) Create(user *user.UserCreate) (string, error) {
	
	query := `INSERT INTO users(id, first_name, last_name, email, identifier, phone, address, city, country, zip_code, password)
						values (?,?,?,?,?,?,?,?,?,?,?)`
	newId := uuid.NewString()
	hashPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return "", err
	}

	err = database.ExecuteTransaction(r.DB, query, newId, user.FirstName, user.LastName, user.Email, user.Identifier, user.Phone, user.Address, user.City, user.Country, user.ZipCode, hashPassword)

	if err != nil {
		return "", err
	}

	return newId, nil	
}

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

func (r *Repository) ExistUser(identifier string, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE identifier = ? OR email = ?)`

	row := database.GetRow(r.DB, query, identifier, email)

	var exist bool
	if err := row.Scan(&exist); err != nil {
		return false, err
	}

	return exist, nil
}