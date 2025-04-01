package routes

import "github.com/gofiber/fiber/v2"

func EntablishmentsRoutes(app *fiber.App) {
	users := app.Group("/entablishments")
	users.Get("/", getEntablishments)       // GET /users
	users.Get("/:id", getEntablishmentById) // GET /users/:id
	users.Post("/", createEntablishment)    // POST /users
}

func getEntablishments(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getEntablishmentById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createEntablishment(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }