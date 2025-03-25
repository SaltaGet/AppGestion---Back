package routes

import "github.com/gofiber/fiber/v2"

func ProductRoutes(app *fiber.App) {
	users := app.Group("/product")
	users.Get("/", getProduct)       // GET /users
	users.Get("/:id", getUProductByID) // GET /users/:id
	users.Post("/", createProduct)    // POST /users
}

func getProduct(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getUProductByID(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createProduct(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }