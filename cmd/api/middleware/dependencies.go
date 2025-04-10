package middleware

import (
	"appGestion/pkg/dependencies"
	"appGestion/pkg/key"
	"context"

	"github.com/gofiber/fiber/v2"
)

func InjectApp(app *dependencies.Application) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext() // Obtiene el context.Context nativo
		ctx = context.WithValue(ctx, key.AppKey, app) // Inyecta dependencias
		c.SetUserContext(ctx) // Vuelve a asignar el contexto a Fiber

		return c.Next()
	}
}
