package auth

import (
	"api-stock/pkg/models/auth"
	resp "api-stock/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (ctrl Controller) Login(c *fiber.Ctx) error {
	id, err := ctrl.AuthService.Login(&auth.AuthLogin{Identifier: "00000000", Password: "Qwer1234*"})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp.Response{
			Status:  false,
			Body:    nil,
			Message: "Error al crear cliente",
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp.Response{
		Status:  true,
		Body:    id,
		Message: id,
	})
}