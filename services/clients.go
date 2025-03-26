package services

import (
	db "api-stock/database"
	m "api-stock/models"
	"api-stock/utils"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
)

func CreateClient(client *m.Client) (int, string, error) {
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

	pass_hash, err := utils.HashPassword(os.Getenv("PASSWORD_ADMIN"))

	if err != nil {
		return fiber.StatusInternalServerError, "Se produjo un error al tratar al cliente", err
	}

	_, err = db.Exec(query,	client.Id, client.Email, client.CUIT, client.Name, pass_hash, client.Cellphone, time.Now(), time.Now(),)

	if err != nil {
		return fiber.StatusInternalServerError, "Se producjo un error al tratar al cliente", err
	}

	return fiber.StatusCreated, client.Id, nil
}

func GetClientByCUIT(id string) (bool, error) {
	var exist bool

	db := db.GetDB()

	query:= `SELECT EXISTS(SELECT 1 FROM clients WHERE cuit = ?)`

	err := db.QueryRow(query, id,).Scan(&exist)

	if err != nil{
		return false, err
	}

	return exist, nil
}