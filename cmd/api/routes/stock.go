package routes

import "github.com/gofiber/fiber/v2"

func StocksRoutes(app *fiber.App) {
	users := app.Group("/stocks")
	users.Get("/", getStocks)       // GET /users
	users.Get("/:id", getStockById) // GET /users/:id
	users.Post("/", createStock)    // POST /users
}

func getStocks(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getStockById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createStock(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }