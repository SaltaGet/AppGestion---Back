package routes

import (
	"github.com/gofiber/fiber/v2"
	entCtrl "api-stock/cmd/api/controllers/entity"
	// mdw "api-stock/cmd/api/middleware"
)

func EntitiesRoutes(app *fiber.App, controller *entCtrl.Controller) {
	entities := app.Group("/entities")
	entities.Get("/", getClients)    
	entities.Get("/:id", getClientById)
	entities.Post("/", controller.CreateEntity) 
}

func getClients(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getClientById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }