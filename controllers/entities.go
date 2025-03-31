package controllers

import (
	cl "api-stock/models/entity"
	m "api-stock/models"
	s "api-stock/services"

	"github.com/gofiber/fiber/v2"
)

func CreateClient(c *fiber.Ctx) error {
	var client cl.ClientCreate

	err := c.BodyParser(&client)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(m.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al crear cliente",
		})
	}

	if err := client.Validate(); err != nil {
		return c.Status(422).JSON(m.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	status, message, err := s.CreateClient(&client)

	if status > 299 || err != nil{
		return c.Status(status).JSON(m.Response{
			Status:  false,
			Body:    nil,
			Message: message,
		})
	}

	return c.Status(status).JSON(m.Response{
		Status:  true,
		Body:    map[string]string{"client_id": message},
		Message: "Cliente creado con éxito",
	})
}

func ClientLogin(c *fiber.Ctx) error {
	var clienLogin cl.ClientLogin

	err := c.BodyParser(&clienLogin)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(m.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al intentar loguearse",
		})
	}

	if err := clienLogin.Validate(); err != nil {
		return c.Status(422).JSON(m.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	status, message, err := s.LoginClient(&clienLogin)

	if status > 299 {
		return c.Status(status).JSON(m.Response{
			Status:  false,
			Body:    err,
			Message: message,
		})
	}

	return c.Status(status).JSON(m.Response{
		Status:  true,
		Body:    map[string]string{"token": message},
		Message: "Token generado con exito",
	})
}