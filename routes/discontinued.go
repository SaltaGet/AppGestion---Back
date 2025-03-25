package routes

import "github.com/gofiber/fiber/v2"

func DiscontinuedRoutes(app *fiber.App) {
	users := app.Group("/discontinued")
	users.Get("/", getDiscontinued)       // GET /users
	users.Get("/:id", getDiscontinuedByID) // GET /users/:id
	users.Post("/", createDiscontinued)    // POST /users
}

func getDiscontinued(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getDiscontinuedByID(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createDiscontinued(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }