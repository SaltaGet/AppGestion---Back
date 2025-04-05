package auth

import (
	"api-stock/pkg/models/auth"
	resp "api-stock/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (ctrl Controller) Login(c *fiber.Ctx) error {
	var authLogin auth.AuthLogin

	err := c.BodyParser(&authLogin)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al realizar login",
		})
	}

	token, code,  err := ctrl.AuthService.Login(&authLogin)

	if err != nil {
		return c.Status(code).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: token,
		})
	}

	return c.Status(code).JSON(resp.Response{
		Status:  true,
		Body:    token,
		Message: "Token obtenido con éxito",
	})
}