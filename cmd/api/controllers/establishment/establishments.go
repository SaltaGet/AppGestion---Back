package establishment

import (
	"github.com/gofiber/fiber/v2"
	"api-stock/pkg/models/establishment"
	resp "api-stock/pkg/models"
)

func (ctrl *Controller) Create(c *fiber.Ctx) error {
	var establishment establishment.EstablishmentCreate
	
	if err := c.BodyParser(&establishment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al leer el JSON de la petición",
		})
	}

	if err := establishment.Validate(); err != nil {
		return c.Status(422).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := ctrl.EstablishmentService.Create(&establishment)

	if err != nil {
		if errResp, ok := err.(*resp.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(resp.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	return c.Status(200).JSON(resp.Response{
		Status:  true,
		Body:    id,
		Message: "Token obtenido con éxito",
	})
}