package establishment

import (
	"api-stock/pkg/models/establishment"
	"api-stock/pkg/repository/database"
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
