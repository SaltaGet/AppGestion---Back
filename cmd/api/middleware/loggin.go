package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now() // ⏱️ Marca el inicio

	err := c.Next() // 📦 Pasa al siguiente middleware / handler

	duration := time.Since(start) // 🕒 Calcula cuánto tardó

	// 🧠 Setea el header de respuesta con la duración
	c.Set("X-Response-Time", duration.String())

	// Opcional: loguearlo
	log.Printf("Request: %s %s took %s", c.Method(), c.Path(), duration)

	return err
}
// func LogginMiddleware(c *fiber.Ctx) error {
// 	log.Printf("Resquest: %s %s", c.Method(), c.Path())

// 	return c.Next()
// }