package controllers

import (
	m "api-stock/models"
	s "api-stock/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateClient(c *fiber.Ctx) error {
	var client m.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(m.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al crear cliente",
		})
	}

	client.Id = uuid.NewString()

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