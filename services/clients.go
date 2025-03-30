package services

import (
	db "api-stock/database"
	cl "api-stock/models/client"
	"api-stock/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"database/sql"
	"os"
	"time"
	"errors"
)

func LoginClient(clientLogin *cl.ClientLogin) (int, string, error) {
	db := db.GetDB()

	query := `SELECT * FROM clients WHERE cuit = ?`

	row := db.QueryRow(query, clientLogin.CUIT)

	var client cl.Client

	err := row.Scan(&client.Id, &client.Email, &client.CUIT, &client.Name,
		&client.Password, &client.Cellphone, &client.Role, &client.IsActive, &client.CreatedAt, &client.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return 404, "Credenciales incorrectas", err
		}
		return 500, "Error al obtener usuario para el login", err
	}

	if !utils.CheckPasswordHash(clientLogin.Password, client.Password) {
		return 404, "Credenciales incorrectas", nil
	}

	token, err := utils.GenerateClientToken(&client)

	if err != nil {
		return 500, "Error al intentar generar el token", err
	}

	return 200, token, nil

}

func CreateClient(client *cl.ClientCreate) (int, string, error) {
	db := db.GetDB()

	exist, err := GetClientByCUIT(client.CUIT)

	if err != nil {
		return fiber.StatusInternalServerError, "Error al comprobar la existencia del cliente", err
	}

	if exist {
		return fiber.StatusBadRequest, "El cliente ya existe", err
	}

	query := `INSERT INTO clients (id, email, cuit, name, password, cellphone, created_at, updated_at)
		VALUES (?,?,?,?,?,?,?,?)`

	newId := uuid.NewString()

	passHash, err := utils.HashPassword(os.Getenv("PASSWORD_ADMIN"))

	if err != nil {
		return fiber.StatusInternalServerError, "Se produjo un error al tratar al cliente", err
	}

	_, err = db.Exec(query, newId, client.Email, client.CUIT, client.Name, passHash, client.Cellphone, time.Now(), time.Now())

	if err != nil {
		return fiber.StatusInternalServerError, "Se producjo un error al tratar al cliente", err
	}

	return fiber.StatusCreated, newId, nil
}

func GetClientByCUIT(id string) (bool, error) {
	var exist bool

	db := db.GetDB()

	query := `SELECT EXISTS(SELECT 1 FROM clients WHERE cuit = ?)`

	err := db.QueryRow(query, id).Scan(&exist)

	if err != nil {
		return false, err
	}

	return exist, nil
}
