package routes

import "github.com/gofiber/fiber/v2"

func SalesRoutes(app *fiber.App) {
	users := app.Group("/sales")
	users.Get("/", getUsers)       // GET /users
	users.Get("/:id", getUserByID) // GET /users/:id
	users.Post("/", createUser)    // POST /users
}

func getSale(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getSalesByID(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createSale(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }