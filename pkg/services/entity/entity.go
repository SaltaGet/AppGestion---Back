package entity

import (
	// db "api-stock/pkg/repository/database"
	ent "api-stock/pkg/models/entity"
	// "api-stock/pkg/utils"
	"github.com/google/uuid"
	// "database/sql"
	// "os"
	// "time"
	// "errors"
)

func (s *Service) LoginClient(clientLogin *ent.ClientLogin) (int, string, error) {
	// query := `SELECT * FROM clients WHERE cuit = ?`

	// row := db.QueryRow(query, clientLogin.CUIT)

	// var client ent.Entity

	// err := row.Scan(&client.Id, &client.Email, &client.CUIT, &client.Name,
	// 	&client.Phone, &client.Role, &client.IsActive, &client.Created, &client.Updated)

	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {

	// 		return 404, "Credenciales incorrectas", err
	// 	}
	// 	return 500, "Error al obtener usuario para el login", err
	// }

	// if !utils.CheckPasswordHash(clientLogin.Password, client.Password) {
	// 	return 404, "Credenciales incorrectas", nil
	// }

	// token, err := utils.GenerateClientToken(&client)

	// if err != nil {
	// 	return 500, "Error al intentar generar el token", err
	// }

	// return 200, token, nil
	return 200, "", nil

}

func (s Service) Create(entity *ent.EntityCreate) (string, error) {
	// exist, err := GetClientByCUIT(entity.CUIT)

	// if err != nil {
	// 	return "", err
	// }

	// if exist {
	// 	return "", err
	// }

	// query := `INSERT INTO clients (id, email, cuit, name, password, phone, created_at, updated_at)
	// 	VALUES (?,?,?,?,?,?,?,?)`

	newId := uuid.NewString()

	// passHash, err := utils.HashPassword(os.Getenv("PASSWORD_ADMIN"))

	// if err != nil {
	// 	return "Se produjo un error al tratar al cliente", err
	// }

	// _, err = db.Exec(query, newId, entity.Email, entity.CUIT, entity.Name, passHash, entity.Cellphone, time.Now(), time.Now())

	// if err != nil {
	// 	return "Se producjo un error al tratar al cliente", err
	// }

	return newId, nil
}

func GetClientByCUIT(id string) (bool, error) {
	// var exist bool

	// db := db.GetDB()

	// query := `SELECT EXISTS(SELECT 1 FROM clients WHERE cuit = ?)`

	// err := db.QueryRow(query, id).Scan(&exist)

	// if err != nil {
	// 	return false, err
	// }

	// return exist, nil
	return true, nil
}

func (s *Service) Insert(entity *ent.EntityCreate) (string, error) {
	return "", nil
}
