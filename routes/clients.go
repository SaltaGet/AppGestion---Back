package routes

import (
	"github.com/gofiber/fiber/v2"
	ctrl "api-stock/controllers"
)

func ClientsRoutes(app *fiber.App) {
	users := app.Group("/clients")
	users.Get("/", getClients)       // GET /users
	users.Get("/:id", getClientById) // GET /users/:id
	users.Post("/", ctrl.CreateClient)    // POST /users
}

func getClients(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getClientById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }