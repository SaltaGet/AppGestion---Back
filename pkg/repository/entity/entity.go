package entity

import (
	"api-stock/pkg/models/entity"
	// "errors"
	// "github.com/google/uuid"
)

func (r *Repository) Insert(entity *entity.EntityCreate) (string, error){
	// query := `INSERT INTO clients (id, email, cuit, name, password, phone, created_at, updated_at)
  //       VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// // Generar un nuevo ID único
	// newId := uuid.NewString()

	// // Ejecutar la consulta para insertar el cliente
	// _, err := r.DB.Exec(query, newId, entity.Email, entity.CUIT, entity.Name, entity.Password, entity.Cellphone, time.Now(), time.Now())
	// if err != nil {
	// 		return "", errors.New("error al insertar el cliente en la base de datos")
	// }

	// // Retornar el ID generado
	// return newId, nil
	return "", nil
}