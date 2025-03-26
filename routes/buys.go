package routes

import "github.com/gofiber/fiber/v2"

func BuysRoutes(app *fiber.App) {
	users := app.Group("/buys")
	users.Get("/", getBuys)       // GET /users
	users.Get("/:id", getBuyById) // GET /users/:id
	users.Post("/", createBuy)    // POST /users
}

func getBuys(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getBuyById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createBuy(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }