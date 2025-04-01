package entity

import (
	ent "api-stock/pkg/models/entity"
	resp "api-stock/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func (ctrl Controller) CreateEntity(c *fiber.Ctx) error {
	var entity ent.EntityCreate

	err := c.BodyParser(&entity)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al crear cliente",
		})
	}

	if err := entity.Validate(); err != nil {
		return c.Status(422).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := ctrl.EntityService.Create(&entity)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al intentar crear entidad",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp.Response{
		Status:  true,
		Body:    map[string]string{"entity_id": id},
		Message: "Cliente creado con éxito",
	})
}

// func ClientLogin(c *fiber.Ctx) error {
// 	var clienLogin cl.ClientLogin

// 	err := c.BodyParser(&clienLogin)

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(m.Response{
// 			Status:  false,
// 			Body:    nil,
// 			Message: "Error al intentar loguearse",
// 		})
// 	}

// 	if err := clienLogin.Validate(); err != nil {
// 		return c.Status(422).JSON(m.Response{
// 			Status:  false,
// 			Body:    nil,
// 			Message: err.Error(),
// 		})
// 	}

// 	status, message, err := s.LoginClient(&clienLogin)

// 	if status > 299 {
// 		return c.Status(status).JSON(m.Response{
// 			Status:  false,
// 			Body:    err,
// 			Message: message,
// 		})
// 	}

// 	return c.Status(status).JSON(m.Response{
// 		Status:  true,
// 		Body:    map[string]string{"token": message},
// 		Message: "Token generado con exito",
// 	})
// }