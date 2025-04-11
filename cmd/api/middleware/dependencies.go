package middleware

import (
	"appGestion/pkg/dependencies"
	"appGestion/pkg/key"
	"context"

	"github.com/gofiber/fiber/v2"
)

func InjectApp(app *dependencies.Application) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext() 
		ctx = context.WithValue(ctx, key.AppKey, app) 
		c.SetUserContext(ctx)

		return c.Next()
	}
}
