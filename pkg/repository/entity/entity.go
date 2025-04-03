package entity

import (
	"api-stock/pkg/models/entity"
	"github.com/google/uuid"
	"errors"
	"time"
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

func (r *Repository) Update(entity *entity.EntityUpdate) error {
	return nil
}