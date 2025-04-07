package auth

import (
	"api-stock/pkg/models/auth"
	resp "api-stock/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// //*********************************************************
// func (ctrl *Controller) GetEntity(c *fiber.Ctx) error {
// 	db := c.Locals("db").(*sql.DB) // Conexión del tenant
	
// 	// Usa la conexión específica
// 	entity, err := ctrl.AuthService.GetEntity(db, c.Params("id"))
// 	if err != nil {
// 		// Manejo de errores
// 	}
	
// 	return c.JSON(entity)
// }
// //*********************************************************

func (ctrl Controller) Login(c *fiber.Ctx) error {
	var authLogin auth.AuthLogin

	if err := c.BodyParser(&authLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al leer el JSON de la petición",
		})
	}

	if err := authLogin.Validate(); err != nil {
		return c.Status(422).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	token, err := ctrl.AuthService.Login(&authLogin)

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
		Body:    token,
		Message: "Token obtenido con éxito",
	})
}