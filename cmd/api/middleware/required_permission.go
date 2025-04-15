package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequiredPermission(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role_tenant").(string)
		if role != permission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "No cuenta con los permisos para ejecutar esta acción",
			})
		}
		return c.Next()
	}
}