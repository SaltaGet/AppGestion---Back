package routes

import "github.com/gofiber/fiber/v2"

func SalesRoutes(app *fiber.App) {
	users := app.Group("/sales")
	users.Get("/", getSales)       // GET /users
	users.Get("/:id", getSalesById) // GET /users/:id
	users.Post("/", createSale)    // POST /users
}

func getSales(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getSalesById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createSale(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }