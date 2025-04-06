package middleware

import (
	"api-stock/pkg/models/user"
	"github.com/gofiber/fiber/v2"
)

func RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(user.User)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "No autorizado",
			})
		}

		// Verificar si el usuario tiene un rol permitido
		for _, role := range allowedRoles {
			if user.Role.Name == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Acceso denegado",
		})
	}
}
