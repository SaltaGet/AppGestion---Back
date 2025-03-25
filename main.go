package main

import (
	"log"
	"api-stock/routes"
	"api-stock/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer database.CloseDB()

	app := fiber.New()

	app.Use(logginMiddleware)

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func logginMiddleware(c *fiber.Ctx) error {
	log.Printf("Resquest: %s %s", c.Method(), c.Path())

	return c.Next()
}