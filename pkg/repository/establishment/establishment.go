package establishment

import (
	"appGestion/pkg/models/establishment"
	"appGestion/pkg/repository/database"
	"appGestion/pkg/utils"

	"github.com/google/uuid"
)

func (r *Repository) Create(establishment *establishment.EstablishmentCreate, connection string) (string, error) {
	query := `INSERT INTO establishments (id, name, email, phone, address, city, country, zip_code, date_create, connection, entity_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	newId := uuid.NewString()

	err := database.ExecuteTransaction(r.DB, query, newId, establishment.Name, establishment.Email, establishment.Phone, establishment.Address,
		establishment.City, establishment.Country, establishment.ZipCode, establishment.DateCreate, connection, establishment.EntityId)

	if err != nil {
		return "", err
	}

	return newId, nil
}

func (r *Repository) GetEstablishmentById(establishmentId string, userId string) (string, error) {
	query := `SELECT connection FROM establishments e
		JOIN establishment_user eu ON e.id = eu.establishment_id
		WHERE e.id = ? AND eu.user_id = ?`

	row := database.GetRow(r.DB, query, establishmentId, userId)

	var connection string
	if err := row.Scan(&connection); err != nil {
		return "", err
	}

	return connection, nil
}

func (r *Repository) GetAllAdmin() (*[]establishment.Establishment, error) {
	query := `SELECT *FROM establishments`

	var establishments []establishment.Establishment
	rows, err := database.GetRows(r.DB, query)

	if err != nil {
		return nil, err
	}

	err = utils.MapRowsToStruct(rows, &establishments)

	if err != nil {
		return nil, err
	}

	return &establishments, nil
}	
