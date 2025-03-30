package routes

import (
	"github.com/gofiber/fiber/v2"
	ctrl "api-stock/controllers"
	mdw "api-stock/middleware"
)

func ClientsRoutes(app *fiber.App) {
	users := app.Group("/clients")
	users.Get("/", getClients)    
	users.Get("/:id", getClientById)
	users.Post("/", mdw.JWTMultiProtected("client"), ctrl.CreateClient) 
	users.Post("/login", ctrl.ClientLogin)   
}

func getClients(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getClientById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }