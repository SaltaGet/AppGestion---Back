package routes

import (
	entCtrl "appGestion/cmd/api/controllers/entity"
	"appGestion/cmd/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func EntitiesRoutes(app *fiber.App, ctrl *entCtrl.Controller) {
	entities := app.Group("/entities")
	entities.Get("/", getClients)    
	entities.Get("/get_all", middleware.RequireAuth(true), ctrl.GetAll)
	entities.Post("/create", middleware.RequireAuth(true), ctrl.CreateEntity) 
}

func getClients(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }