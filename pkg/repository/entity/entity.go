package entity

import (
	"appGestion/pkg/models/entity"
	"appGestion/pkg/repository/database"
	"appGestion/pkg/utils"

	"github.com/google/uuid"
)

func (r *Repository) Insert(entity *entity.EntityCreate) (string, error){
	query := `INSERT INTO entities (id, email, cuit, name, phone, start_activities, address, city, country, zip_code)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	newId := uuid.NewString()

	err := database.ExecuteTransaction(r.DB, query,newId, entity.Email, entity.CUIT, entity.Name, entity.Phone, entity.StartActitivies, entity.Address, entity.City, entity.Country, entity.ZipCode)	

	if err != nil {
			return "", err
	}

	return newId, nil
}

func (r *Repository) ExistById(id string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM entities WHERE id = ?)`

	row := database.GetRow(r.DB, query, id)

	var exist bool
	if err := row.Scan(&exist); err != nil {
		return false, err
	}

	return exist, nil
}

func (r *Repository) ExistByCUIT(cuit string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM entities WHERE cuit = ?)`

	row := database.GetRow(r.DB, query, cuit)

	var exist bool
	if err := row.Scan(&exist); err != nil {
		return false, err
	}

	return exist, nil
}

func (r *Repository) GetAll() (*[]entity.Entity, error) {
	query := `SELECT * FROM entities`

	rows, err := database.GetRows(r.DB, query)

	if err != nil {
		return nil, err
	}

	var entities []entity.Entity
	err = utils.MapRowsToStruct(rows, &entities)

	if err != nil {
		return nil, err
	}

	return &entities, nil 
}

func (r *Repository) Update(entity *entity.EntityUpdate) error {
	return nil
}