package routes

import "github.com/gofiber/fiber/v2"

func CategoriesRoutes(app *fiber.App) {
	users := app.Group("/categories")
	users.Get("/", getCategories)       // GET /users
	users.Get("/:id", getCategoryById) // GET /users/:id
	users.Post("/", createCategory)    // POST /users
}

func getCategories(c *fiber.Ctx) error       { return c.SendString("Lista de usuarios") }
func getCategoryById(c *fiber.Ctx) error    { return c.SendString("Usuario ID: " + c.Params("id")) }
func createCategory(c *fiber.Ctx) error     { return c.SendString("Usuario creado") }