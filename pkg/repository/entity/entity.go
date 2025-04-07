package entity

import (
	"api-stock/pkg/models/entity"
	"api-stock/pkg/repository/database"
	"errors"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) Insert(entity *entity.EntityCreate) (string, error){
	query := `INSERT INTO entities (id, email, cuit, name, password, phone, created, updated)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Generar un nuevo ID único
	newId := uuid.NewString()

	// Ejecutar la consulta para insertar el cliente
	_, err := r.DB.Exec(query, newId, entity.Email, entity.CUIT, entity.Name, entity.Password, entity.Cellphone, time.Now().UTC(), time.Now().UTC())
	if err != nil {
			return "", errors.New("error al insertar el cliente en la base de datos")
	}

	// Retornar el ID generado
	return newId, nil
}

func (r *Repository) Exist(id string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM entities WHERE id = ?)`

	row := database.GetRow(r.DB, query, id)

	var exist bool
	if err := row.Scan(&exist); err != nil {
		return false, err
	}

	return exist, nil
}

func (r *Repository) Update(entity *entity.EntityUpdate) error {
	return nil
}