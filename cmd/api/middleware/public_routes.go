// middleware/public.go
package middleware

import "github.com/gofiber/fiber/v2"

func PublicRoutes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Marcar la ruta como pública
		c.Locals("is_public", true)
		return c.Next()
	}
}