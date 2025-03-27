package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func LogginMiddleware(c *fiber.Ctx) error {
	log.Printf("Resquest: %s %s", c.Method(), c.Path())

	return c.Next()
}