package routes

import (
	"appGestion/cmd/api/controllers/establishment"
	"appGestion/cmd/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func EntablishmentsRoutes(app *fiber.App, ctrl *establishment.Controller) {
	establishments := app.Group("/establishments")
	establishments.Get("/", getEntablishments)       // GET /users
	establishments.Get("/:id", getEntablishmentById) // GET /users/:id
	establishments.Post("/create", middleware.RequireAuth(true), ctrl.Create)    // POST /users
}

func getEntablishments(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getEntablishmentById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
